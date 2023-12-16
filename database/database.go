package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Dbase *sql.DB

func StartDB() {
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping db!")
	}

	_, err1 := db.Exec(ReadTable("database.sql"))
	if err1 != nil {
		log.Fatal(err1)
	}
	// _, err2 := db.Exec(ReadTable("ThreadsPopulate.sql"))
	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	Dbase = db
}

func ReadTable(filename string) string {
	data, _ := os.ReadFile("./db/" + filename)
	return string(data)
}
