package models

import (
	db "background/database"
	"log"
)

type Person struct {
	Id string `json:"id" form:"id"`
	Password string `json:"password"`
	Name string `json:"name" form:"name"`
}

// GetPeople 返回1代表学生，0代表老师，-1表示账号或密码错误
func (p *Person) GetPeople() (existed int){
	err := db.SqlDB.QueryRow("SELECT IFNULL((SELECT 1 FROM student WHERE id = ? AND password = ? LIMIT 1), IFNULL((SELECT 0 FROM teacher WHERE id = ? AND password = ? LIMIT 1), -1)) AS existed", p.Id, p.Password, p.Id, p.Password).Scan(&existed)
	if err != nil {
		log.Fatal(err)
	}
	return
}
