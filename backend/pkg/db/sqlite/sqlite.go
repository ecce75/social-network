package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func createURL(path, prefix string) string {
	urlPath := strings.Replace(path, "\\", "/", -1)
	if runtime.GOOS == "windows" && strings.Contains(urlPath, ":/") {
		urlPath = "/" + urlPath
	}
	return prefix + "://" + urlPath
}

func ConnectAndMigrate(dbPath string, migrationsPath string) (*sql.DB, error) {
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

	migrationsURL := createURL(migrationsPath, "file")
	dbURL := createURL(dbPath, "sqlite")

	fmt.Printf("Applying migrations from path: %s\n", migrationsURL)

	m, err := migrate.New(migrationsURL, dbURL)
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
	return db, nil
}
