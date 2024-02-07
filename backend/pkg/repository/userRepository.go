package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// # Data access layer, interacts with db

func (r *UserRepository) GetUserByEmailOrNickname(emailOrNickname string) (model.User, error) {
	query := "SELECT * FROM users WHERE email = ? OR username = ? LIMIT 1"
	var user model.User
	err := r.db.QueryRow(query, emailOrNickname, emailOrNickname).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, 
		&user.DOB, &user.AvatarURL, &user.About, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found in database")
		}
		return model.User{}, err
	}
	return user, nil
}

func (r *UserRepository) RegisterUser(data model.RegistrationData) (int64, error) {
	result, err := r.db.Exec("INSERT INTO users (username, email, password, first_name, last_name, date_of_birth, avatar_url, about_me) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
	data.Username, data.Email, data.Password, data.FirstName, data.LastName, data.DOB, data.AvatarURL, data.About)
	if err != nil {
		fmt.Println("Error inserting user into database")
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last inserted user id")
		return 0, err
	}
	return lastInsertID, nil
}