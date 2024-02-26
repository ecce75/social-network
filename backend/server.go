package main

import (
	"backend/api"
	"backend/pkg/db/sqlite"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	dbPath := "./pkg/db/database.db"
	migrationsPath := "pkg/db/migrations/sqlite"

	// Connect to the database and apply migrations
	db, err := sqlite.ConnectAndMigrate(dbPath, migrationsPath)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	api.Router(mux, db)

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	address := os.Getenv("BACKEND_URL")
	if address == "" {
		address = "localhost"
	}

	fmt.Println("Server is running on " + address + ":" + port)
	http.ListenAndServe(":" + port, nil)
}
