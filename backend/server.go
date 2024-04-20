package main

import (
	"backend/api"
	"backend/pkg/db/sqlite"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	mux := mux.NewRouter()
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable not set")
	}
	migrationsPath := "pkg/db/migrations/sqlite"

	// Connect to the database and apply migrations
	db, err := sqlite.ConnectAndMigrate(dbPath, migrationsPath)
	if err != nil {
		fmt.Printf("Failed to connect to the database: %s\n", err)
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()
	// ENV variables

	// TODO: remove in prod
	// err = godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	port := os.Getenv("NEXT_PUBLIC_BACKEND_PORT")
	if port == "" {
		port = "empty"
	}
	address := os.Getenv("NEXT_PUBLIC_URL")
	if address == "" {
		address = "empty"
	}

	// Start the router
	api.Router(mux, db)

	fmt.Println("Server is running on " + address + ":" + port)
	http.ListenAndServe(":"+port, nil)
}
