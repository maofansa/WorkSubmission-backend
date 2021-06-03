package main

import (
	db "background/database"
	"database/sql"
	"log"
)

func main() {
	defer func(SqlDB *sql.DB) {
		err := SqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db.SqlDB)
	router := initRouter()
	err := router.Run(":9190")
	if err != nil {
		log.Fatal(err)
	}
}
