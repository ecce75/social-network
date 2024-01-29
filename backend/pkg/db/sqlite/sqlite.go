package sqlite

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/sqlite"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "log"
)

func ConnectAndMigrate(dbPath string, migrationsPath string) (*sql.DB, error) {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }

    // Run migrations
    m, err := migrate.New(
        "file://"+migrationsPath, // File path for migrations
        "sqlite://"+dbPath,       // Database URL
    )
    if err != nil {
        return nil, err
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return nil, err
    }

    return db, nil
}
