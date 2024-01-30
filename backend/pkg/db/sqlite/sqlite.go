package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var Dbase *sql.DB

func ConnectAndMigrate(dbPath string, migrationsPath string) (*sql.DB, error) {
    // Adjusted for clarity in logging
    fmt.Printf("Connecting to SQLite database at path: %s\n", dbPath)

    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        fmt.Println("Could not connect to SQLite database:", err)
        return nil, err
    }

    err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping db!")
	}

    fmt.Println("Connected to SQLite database successfully.")

    // Ensure migrationsPath is also correctly set relative to your working directory
    fmt.Printf("Applying migrations from path: %s\n", migrationsPath)

    m, err := migrate.New(
        "file://"+migrationsPath,
        "sqlite://"+dbPath,
    )
    if err != nil {
        fmt.Println("Failed to prepare migrations:", err)
        return nil, err
    }

    err = m.Up()
    if err != nil && err != migrate.ErrNoChange {
        fmt.Println("Failed to apply migrations:", err)
        return nil, err
    }

    fmt.Println("Migrations applied successfully.")

    Dbase = db
    return db, nil
}