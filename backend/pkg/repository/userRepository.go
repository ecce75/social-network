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
		&user.DOB, &user.AvatarURL, &user.About, &user.Profile, &user.CreatedAt, &user.UpdatedAt)
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

func (r *UserRepository) GetUserProfileByID(id int) (model.Profile, error) {
	query := "SELECT id, username, first_name, last_name, date_of_birth, avatar_url, about_me, profile, created_at FROM users WHERE id = ?"
	var profile model.Profile
	err := r.db.QueryRow(query, id).Scan(&profile.Id, &profile.Username, &profile.FirstName, &profile.LastName, &profile.DOB, &profile.AvatarURL, &profile.About, &profile.ProfileSetting, &profile.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found in database")
		}
		return model.Profile{}, err
	}
	return profile, nil
}

func (r *UserRepository) UpdateUserProfile(id int, data model.RegistrationData) error {
	_, err := r.db.Exec("UPDATE users SET username = ?, email = ?, password = ?, first_name = ?, last_name = ?, date_of_birth = ?, avatar_url = ?, about_me = ?, profile = ? WHERE id = ?",
		data.Username, data.Email, data.Password, data.FirstName, data.LastName, data.DOB, data.AvatarURL, data.About, data.ProfileSetting, id)
	if err != nil {
		fmt.Println("Error updating user profile in database")
		return err
	}
	return nil
}

func (r *UserRepository) GetAllUsersExcludeRequestingUserAndFriends(userID int) ([]model.UserList, error) {
	query := `
    SELECT u.id, u.username, u.first_name, u.last_name, u.avatar_url 
    FROM users u
    WHERE u.id != ?
    AND u.id NOT IN (
        SELECT f.user_id2 FROM friends f WHERE f.user_id1 = ? AND f.status = 'accepted'
        UNION
        SELECT f.user_id1 FROM friends f WHERE f.user_id2 = ? AND f.status = 'accepted'
    )`

	rows, err := r.db.Query(query, userID, userID, userID)
	if err != nil {
		fmt.Println("Error getting all users excluding friends from database:", err)
		return nil, err
	}
	defer rows.Close()

	var users []model.UserList
	for rows.Next() {
		var user model.UserList
		if err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.AvatarURL); err != nil {
			fmt.Println("Error scanning user data from database:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
