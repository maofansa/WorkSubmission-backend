package models

import (
	db "background/database"
	"database/sql"
	"log"
)

type Teacher struct {
	Id string `json:"id" form:"id"`
	Password string `json:"password"`
	Name string `json:"name" form:"name"`
}

func (t *Teacher) AddTeacher() {

}

func (t *Teacher) GetClassesTeacher() []Class{
	classes := make([]Class, 0)
	rows, err := db.SqlDB.Query("SELECT id, name, picURL FROM class WHERE teacher_id = ?", t.Id)
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
		/*
		if len(class.PicURL) == 0 {
			class.PicURL = "/res/classes/default.png"
		}

		 */
		classes = append(classes, class)
	}
	return classes
}
