package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
)

// # Data access layer, interacts with db

func GetUserByEmailOrNickname(db *sql.DB, emailOrNickname string) (model.User, error) {
	query := "SELECT * FROM users WHERE email = ? OR username = ? LIMIT 1"
	var user model.User
	err := db.QueryRow(query, emailOrNickname, emailOrNickname).Scan(
		&user.UserID, &user.Email, &user.Password, &user.FirstName, &user.LastName, 
		&user.DOB, &user.AvatarURL, &user.Nickname, &user.About)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found in database")
		}
		return model.User{}, err
	}
	return user, nil
}