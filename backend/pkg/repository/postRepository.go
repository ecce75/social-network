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


// GetAllPostsWithUserIDAccess retrieves all posts with the given user ID access.
// It queries the database to fetch posts that meet the following conditions:
// - Posts with the specified user ID
// - Posts with privacy setting set to 'public'
// - Posts with privacy setting set to 'private' and the user is a friend (status = 'accepted')
// The function returns a slice of model.Post and an error if any occurred during the query.
func GetAllPostsWithUserIDAccess(db *sql.DB, userID int) ([]model.Post, error) {
    query := `
    SELECT posts.* 
    FROM posts 
    WHERE posts.user_id = ? 
    OR posts.privacy_setting = 'public' 
    OR (posts.privacy_setting = 'private' AND posts.user_id IN (
        SELECT user_id1 FROM friends WHERE user_id2 = ? AND status = 'accepted'
        UNION
        SELECT user_id2 FROM friends WHERE user_id1 = ? AND status = 'accepted'
    ))
    `

    rows, err := db.Query(query, userID, userID, userID)
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
