package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
	"os"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) GetCommentsByUserID(id int) ([]model.Comment, error) {
	query := `SELECT * FROM comments WHERE user_id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.Id, &comment.PostID, &comment.UserID, &comment.Content, &comment.Image, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) CreateComment(comment *model.Comment) (int64, error) {
	query := `INSERT INTO comments (post_id, user_id, content) 
	VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, comment.PostID, comment.UserID, comment.Content)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if comment.Image.Valid && comment.Image.String == "" {
		return lastInsertID, nil
	}
	comment.Image.String = os.Getenv("NEXT_PUBLIC_URL") + ":" + os.Getenv("NEXT_PUBLIC_BACKEND_PORT") + "/images/comments/" + fmt.Sprint(lastInsertID) + ".jpg"
	r.AddImageUrlToComment(int(lastInsertID), comment.Image.String)
	if err != nil {
		fmt.Println("Error getting last inserted comment id")
	}
	return lastInsertID, nil
}

func (r *CommentRepository) AddImageUrlToComment(postID int, imageURL string) error {
	query := `UPDATE comments SET image_url = ? WHERE id = ?`
	_, err := r.db.Exec(query, imageURL, postID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) GetCommentImageURL(id int) (string, error) {
	query := `SELECT image_url FROM comments WHERE id = ?`
	var imageURL sql.NullString
	err := r.db.QueryRow(query, id).Scan(&imageURL)
	if err != nil {
		return "", err
	}
	if imageURL.Valid {
		return imageURL.String, nil
	}
	return "", nil
}

func (r *CommentRepository) DeleteComment(id int, userid int) error {
	query := `DELETE FROM comments WHERE id = ? AND user_id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) UpdateComment(commentId int, userId int, comment model.UpdateCommentRequest) error {
	query := `UPDATE comments SET content = ? WHERE id = ? AND user_id = ?`
	_, err := r.db.Exec(query, comment.Content, comment.Id, comment.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) GetAllPostComments(id int) ([]model.Comment, error) {
	query := `SELECT * FROM comments WHERE post_id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.Id, &comment.PostID, &comment.UserID, &comment.Content, &comment.Image, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
