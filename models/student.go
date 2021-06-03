package models

import (
	db "background/database"
	"database/sql"
	"log"
)

type Student struct {
	Id string `json:"id" form:"id"`
	Password string `json:"password"`
	Name string `json:"name" form:"name"`
	ClassId []string `json:"class_id" form:"class_id"`
}

func (s *Student) AddStudent() {
	_, err := db.SqlDB.Exec("INSERT INTO student(id, password) VALUES (?, ?)", s.Id, s.Password)
	if err != nil {
		return
	}
}

func (s *Student) GetClassesStudent() []Class{
	classes := make([]Class, 0)
	rows, err := db.SqlDB.Query("SELECT c.id, c.name, c.picURL from class c INNER JOIN student_class s ON s.class_id = c.id AND s.student_id = ?", s.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var class Class
		err := rows.Scan(&class.Id, &class.Name, &class.PicURL)
		if err != nil {
			log.Fatal(err)
		}
		if len(class.PicURL) == 0 {
			class.PicURL = "/res/classes/default.png"
		}
		classes = append(classes, class)
	}
	return classes
}

func (s *Student) JoinClass(classId string) (joined int) {
	// 记录已存在
	var existed int
	err := db.SqlDB.QueryRow("SELECT IFNULL((SELECT 1 FROM student_class WHERE student_id = ? AND class_id = ? LIMIT 1), 0) AS joined", s.Id, classId).Scan(&existed)
	if err != nil {
		log.Fatal(err)
	}
	if existed == 1 {
		return 0
	}

	res, err := db.SqlDB.Exec("INSERT INTO student_class SELECT s.id, c.id FROM student s JOIN class c ON s.id = ? AND c.id = ?", s.Id, classId)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if rowCnt > 0 {
		joined = 1
	}
	return
}

func (s *Student) GetHomeworkListStudent(classId string, rank int) (hwList []Homework) {
	var rows *sql.Rows
	var err error
	if rank == 0 {
		rows, err = db.SqlDB.Query("SELECT i.pid, i.name, i.start_time, i.deadline, IFNULL(h.url, '0') AS url FROM (SELECT * FROM homework_info WHERE class_id = ?) AS i LEFT JOIN (SELECT * FROM homework WHERE student_id = ?) AS h ON i.pid = h.pid ORDER BY (i.start_time)", classId, s.Id)
	} else {
		rows, err = db.SqlDB.Query("SELECT i.pid, i.name, i.start_time, i.deadline, IFNULL(h.url, '0') AS url FROM (SELECT * FROM homework_info WHERE class_id = ?) AS i LEFT JOIN (SELECT * FROM homework WHERE student_id = ?) AS h ON i.pid = h.pid ORDER BY h.url, i.start_time", classId, s.Id)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		var hw Homework
		err := rows.Scan(&hw.Info.Pid, &hw.Info.Name, &hw.Info.StartTime, &hw.Info.Deadline, &hw.URL)
		if hw.URL != "0" {
			hw.Handled = 1
		}
		if err != nil {
			log.Fatal(err)
		}
		hwList = append(hwList, hw)
	}
	return
}