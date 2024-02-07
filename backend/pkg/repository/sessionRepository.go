package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) StoreSessionInDB(sessionToken string, userID int) {
	expiresAt := time.Now().Add(15*time.Minute)
	_, err := r.db.Exec(`INSERT OR REPLACE INTO sessions (sessionToken, userID, expiresAt)
	VALUES (?, ?, ?)`, sessionToken, userID, expiresAt)
	if err != nil {
		fmt.Println("Error inserting session into database: ", err)
		log.Fatalf("Error inserting session into database: %v", err)
	}
}

func (r *SessionRepository) GetSessionBySessionToken(sessionToken string) (model.Session, error){
	var session model.Session
	err := r.db.QueryRow(`SELECT userID, expiresAt FROM sessions WHERE sessionToken = ?`, sessionToken).Scan(&session.UserID, &session.ExpiresAt)
	if err != nil {
		fmt.Println("Error querying session: ", err)
		return model.Session{}, err
	}
	return session, nil
}

func (r *SessionRepository) GetUserIDFromSessionToken(sessionToken string) (int, error) {
	var userID int
	err := r.db.QueryRow(`SELECT userID FROM sessions WHERE sessionToken = ?`, sessionToken).Scan(&userID)
	if err != nil {
		fmt.Println("Error querying session: ", err)
		return 0, err
	}
	return userID, nil
}