package backend

import (
	"fmt"
	"log"
	"net/http"
	"/pkg/db/sqlite"
)

func main() {

	dbPath := "path/to/your/database/file.sqlite"
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