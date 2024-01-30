package model

// Data structures and domain model

type User struct {
	UserID 		int
	Email 		string
	Password 	string
	FirstName 	string
	LastName 	string
	DOB 		string
	AvatarURL 	string
	Nickname 	string
	About 		string
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}