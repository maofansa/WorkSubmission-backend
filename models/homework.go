package models

import (
	db "background/database"
	"database/sql"
	"log"
)

type HomeworkInfo struct {
	Pid string `json:"id"`
	ClassId string `json:"class_id"`
	TeacherId string `json:"teacher_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	StartTime string `json:"pub_time"`
	Deadline string `json:"ddl"`
	ClassNum int `json:"class_num"`
	HandledNum int `json:"handled_num"`
}

type Homework struct {
	Info HomeworkInfo `json:"info"`
	Handled int `json:"status"`
	StudentName string `json:"student_name"`
	StudentId string `json:"student_id"`
	Filename string `json:"filename"`
	Score float64 `json:"score"`
	Commit string `json:"commit"`
	URL string `json:"url"`
	FileName string `json:"file_name"`
}

func (info *HomeworkInfo) ExistedHomeworkInfo() bool{
	var existed int
	err := db.SqlDB.QueryRow("SELECT IFNULL((SELECT 1 FROM homework_info WHERE pid = ? AND class_id = ?), 0) AS existed", info.Pid, info.ClassId).Scan(&existed)
	if err != nil {
		log.Fatal(err)
	}
	if existed == 1 {
		return true
	}
	return false
}

func (h *Homework) UploadHomework() int {
	var existed int
	err := db.SqlDB.QueryRow("SELECT IFNULL((SELECT 1 FROM homework WHERE pid = ? AND class_id = ? AND student_id = ?), 0) AS existed;", h.Info.Pid, h.Info.ClassId, h.StudentId).Scan(&existed)
	if err != nil {
		return 403
	}

	if existed == 1 {
		_, err := db.SqlDB.Exec("UPDATE homework SET url = ? WHERE pid = ? AND student_id =? AND class_id = ?", h.URL, h.Info.Pid, h.StudentId, h.Info.ClassId)
		if err != nil {
			return 404
		}
	} else {
		_, err := db.SqlDB.Exec("INSERT INTO homework(pid, student_id, class_id, url, filename) VALUES (?, ?, ?, ?, ?)", h.Info.Pid, h.StudentId, h.Info.ClassId, h.URL, h.Filename)
		if err != nil {
			return 405
		}
	}
	return 200
}

func (info *HomeworkInfo) CreateHomeworkInfo() (int, error) {
	_, err := db.SqlDB.Exec("INSERT INTO homework_info(pid, name, start_time, deadline, class_id) VALUES (UUID(), ?, STR_TO_DATE(?, '%Y/%m/%d/%H:%i'), STR_TO_DATE(?, '%Y/%m/%d/%H:%i'), ?)", info.Name, info.StartTime, info.Deadline, info.ClassId)
	if err != nil {
		return 0, err
	}
	return 1, err
}

func (h *Homework) DownloadHomework() (string, string, error) {
	var fileURL string
	var fileName string
	err := db.SqlDB.QueryRow("SELECT url, filename FROM homework WHERE class_id = ? AND pid = ? AND student_id = ?", h.Info.ClassId, h.Info.Pid, h.StudentId).Scan(&fileURL, &fileName)
	if err != nil {
		//log.Fatal(err)
	}
	return fileURL, fileName, err
}

func (info *HomeworkInfo) DownloadAllHomework() []Homework{
	homeworks := make([]Homework, 0)
	rows, err := db.SqlDB.Query("SELECT url, filename FROM homework WHERE class_id = ? AND pid = ?", info.ClassId, info.Pid)

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
		var h Homework
		err := rows.Scan(&h.URL, &h.Filename)
		if err != nil {
			log.Fatal(err)
		}
		homeworks = append(homeworks, h)
	}
	return homeworks
}

func (info *HomeworkInfo) GetHomeworkList() []Homework {
	homeworks := make([]Homework, 0)
	rows, err := db.SqlDB.Query("SELECT s.name, h.student_id, IFNULL(h.score, 0) AS grade, IFNULL((SELECT 1 FROM homework h2 WHERE h2.pid = ? AND h2.class_id = ? AND h2.student_id = h.student_id), 0) AS handled FROM homework h INNER JOIN student s on h.student_id = s.id AND h.class_id = ? AND h.pid = ?", info.Pid, info.ClassId, info.Pid, info.ClassId)
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
		var h Homework
		err := rows.Scan(&h.StudentName, &h.StudentId, &h.Score, &h.Handled)
		if err != nil {
			log.Fatal(err)
		}
		homeworks = append(homeworks, h)
	}
	return homeworks
}

func (h *Homework) AddScore() int {
	_, err := db.SqlDB.Exec("UPDATE homework SET score = ? WHERE student_id = ? AND class_id = ? AND pid = ?", h.Score, h.StudentId, h.Info.ClassId, h.Info.Pid)
	if err != nil {
		return 0
	}
	/*
	var score int
	err = db.SqlDB.QueryRow("SELECT score FROM homework WHERE student_id = ? AND class_id = ? AND pid = ?", h.StudentId, h.Info.ClassId, h.Info.Pid).Scan(&score)

	if err != nil {
		return 0
	}

	 */
	return 1
}