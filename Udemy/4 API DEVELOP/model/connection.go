package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var con *sql.DB

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/g_test")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database")
	con = db
	return db
}
