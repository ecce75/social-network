package model

import "time"

// Data structures and domain model

var UserID int

type User struct {
	Id 			int
	Username 	string
	Email 		string
	Password 	string
	FirstName 	string
	LastName 	string
	DOB 		string
	AvatarURL 	string
	About 		string
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationData struct {
	Username 	string `json:"username"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	DOB 		string `json:"dob"`
	AvatarURL 	string `json:"avatar_url,omitempty"`
	About 		string `json:"about"`
}

type AuthResponse struct {
	IsAuthenticated bool `json:"is_authenticated"`
}

type Session struct {
	Id 				int 		`json:"id"`
	SessionToken 	string 		`json:"session_token"`
	UserID 			int 		`json:"user_id"`
	ExpiresAt 		time.Time 	`json:"expires_at"`
}

type Post struct {
	Id 				int 		`json:"id"`
	UserID 			int 		`json:"user_id"`
	Title			string 		`json:"title"`
	Content 		string 		`json:"content,omitempty"`
	ImageURL 		string 		`json:"image_url,omitempty"`
	PrivacySetting 	string    	`json:"privacy_setting"`
    CreatedAt      	time.Time 	`json:"created_at"`
}

type CreatePostRequest struct {
	Title 			string `json:"title"`
	Content 		string `json:"content,omitempty"`
	ImageURL 		string `json:"image_url,omitempty"`
	PrivacySetting 	string `json:"privacy_setting"`
}