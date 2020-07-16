package database

import (
	"database/sql"
	"fmt"
	"log"
)


func Connect(databasename string) *sql.DB {
	db, err := sql.Open("mysql", "root:pass@(localhost:3306)/"+databasename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")
	return db
}