package models

import (
	db "background/database"
	"database/sql"
	"log"
)

type Class struct {
	Id string `json:"id"`
	Name string `json:"name"`
	PicURL string `json:"pic_url"`
	TeacherId string `json:"teacher_id"`
}

func (c *Class) CreateClass() int {
	_, err := db.SqlDB.Exec("INSERT INTO class(ID, NAME, TEACHER_ID, PICURL) VALUES (UUID(), ?, ?, ?)", c.Name, c.TeacherId, c.PicURL)
	if err != nil {
		return 400
	}
	return 200
}

func (c *Class) GetHomeworkListTeacher(rank int) []HomeworkInfo{
	hwList := make([]HomeworkInfo, 0)
	var rows *sql.Rows
	var err error
	var classNum int

	err = db.SqlDB.QueryRow("SELECT COUNT(*) FROM student_class WHERE class_id = ?", c.Id).Scan(&classNum)
	if err != nil {
		log.Fatal(err)
	}

	if rank == 0 {
		rows, err = db.SqlDB.Query("SELECT i.pid, i.name, i.start_time, i.deadline, IFNULL((SELECT COUNT(*) FROM homework WHERE class_id = ? AND pid = i.pid), 0) AS handlednum FROM homework_info i LEFT JOIN homework h on i.pid = h.pid WHERE i.class_id = ? ORDER BY i.start_time", c.Id, c.Id)
	} else {
		rows, err = db.SqlDB.Query("SELECT i.pid, i.name, i.start_time, i.deadline, IFNULL((SELECT COUNT(*) FROM homework WHERE class_id = ? AND pid = i.pid), 0) AS handlednum FROM homework_info i LEFT JOIN homework h on i.pid = h.pid WHERE i.class_id = ? ORDER BY handlednum, i.start_time", c.Id, c.Id)
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
		var info HomeworkInfo
		err := rows.Scan(&info.Pid, &info.Name, &info.StartTime, &info.Deadline, &info.HandledNum)
		if err != nil {
			log.Fatal(err)
		}
		info.ClassNum = classNum
		hwList = append(hwList, info)
	}
	return hwList
}

func (c *Class) GetStudentList() []Student {
	studentList := make([]Student, 0)
	rows, err := db.SqlDB.Query("SELECT s.Name, s.Id FROM student_class c INNER JOIN student s on c.student_id = s.id AND c.class_id = ?", c.Id)
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
		var s Student
		err := rows.Scan(&s.Name, &s.Id)
		if err != nil {
			log.Fatal(err)
		}
		studentList = append(studentList, s)
	}
	return studentList
}

func (c *Class) DeleteStudent(studentId string) (deleted int) {
	_, err := db.SqlDB.Exec("DELETE FROM student_class WHERE class_id = ? AND student_id = ?", c.Id, studentId)
	if err != nil {
		return
	}
	return 1
}

/*
func (c *Class) DeleteClass() int{
	_, err := db.SqlDB.Exec("DELETE FROM student_class WHERE class_id = ?", c.Id)
	_, err := db.SqlDB.Exec("DELETE FROM homework_info WHERE class_id = ?", c.Id)
	_, err := db.SqlDB.Exec("DELETE FROM class WHERE class_id = ?", c.Id)
}
*/
