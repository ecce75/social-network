package main

import (
	"backend/pkg/db/sqlite"
	"fmt"
	"log"
	"net/http"
	//"os"
)

func main() {
    // //working directory check
    // wd, err := os.Getwd()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.Println("Current working directory:", wd)
    // //check end

	dbPath := "./pkg/db/database.db"
    migrationsPath := "pkg/db/migrations/sqlite"

    // Connect to the database and apply migrations
    db, err := sqlite.ConnectAndMigrate(dbPath, migrationsPath)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    defer db.Close()


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world!")
    })

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}