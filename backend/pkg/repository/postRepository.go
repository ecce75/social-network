package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
)

func CreatePost(db *sql.DB, post model.CreatePostRequest, userID int) (int64, error) {
	query := `INSERT INTO posts (user_id, title, content, image_url, privacy_setting) 
	VALUES (?, ?, ?, ?, ?)`
	result, err := db.Exec(query, userID, post.Title, post.Content, post.ImageURL, post.PrivacySetting)
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

func GetAllPostsWithUserIDAccess(db *sql.DB, userID int) ([]model.Post, error) {
	// TODO: implement friends relationship check and private posts
	query := `SELECT * FROM posts WHERE user_id = ? OR privacy_setting = 'public'`

	rows, err := db.Query(query, userID)
	if err != nil {
		return []model.Post{}, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.Id, &post.UserID, &post.Title, &post.Content, &post.ImageURL, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func DeletePost(db *sql.DB, postID int, userID int) error {
	query := `DELETE FROM posts WHERE id = ? AND user_id = ?`
	result, err := db.Exec(query, postID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no post found with the specified id that belongs to the user")
	}
	return nil
}

func UpdatePost(db *sql.DB, postID int, userID int, request model.UpdatePostRequest) error {
    query := `UPDATE posts SET title = ?, content = ?, image_url = ?, privacy_setting = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ? AND user_id = ?`

    result, err := db.Exec(query, request.Title, request.Content, request.ImageURL, request.PrivacySetting, postID, userID)
    if err != nil {
        return err // Handle the error appropriately
    }

    // Check if a row was actually updated
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err // Handle the error appropriately
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no post found with the specified id that belongs to the user or no update was needed")
    }

    return nil
}
