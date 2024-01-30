package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func StoreSessionInDB(db *sql.DB, sessionToken string, userID int) {
	expiresAt := time.Now().Add(15*time.Minute)
	_, err := db.Exec(`INSERT INTO sessions (sessionToken, userID, expiresAt)
	VALUES (?, ?, ?)`, sessionToken, userID, expiresAt)
	if err != nil {
		fmt.Println("Error inserting session into database: ", err)
		log.Fatalf("Error inserting session into database: %v", err)
	}
}

func GetSessionBySessionToken(db *sql.DB, sessionToken string) (model.Session, error){
	var session model.Session
	err := db.QueryRow(`SELECT userID, expiresAt FROM sessions WHERE sessionToken = ?`, sessionToken).Scan(&session.UserID, &session.ExpiresAt)
	if err != nil {
		fmt.Println("Error querying session: ", err)
		return model.Session{}, err
	}
	return session, nil
}

	
