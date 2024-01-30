package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
)

func CreatePost(db *sql.DB, post model.CreatePostRequest, UserID int) (int64, error) {
	query := `INSERT INTO posts (user_id, title, content, image_url, privacy_setting) 
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.Exec(query, UserID, post.Title, post.Content, post.ImageURL, post.PrivacySetting)
	if err != nil {
		fmt.Println("Error inserting post into database: ", err)
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last inserted post id")
	}
	return lastInsertID, nil
}