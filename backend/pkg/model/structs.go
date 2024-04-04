package model

import (
	"database/sql"
	"time"
)

// Data structures and domain model

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	DOB       string
	AvatarURL string
	About     string
	Profile   string
	CreatedAt string
	UpdatedAt string
}

type UserList struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`
}

type Attendance struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Status    string `json:"status"`
}

type Profile struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DOB            string `json:"dob"`
	AvatarURL      string `json:"avatar_url"`
	About          string `json:"about"`
	ProfileSetting string `json:"profile_setting"`
	CreatedAt      string `json:"created_at"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationData struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DOB            string `json:"dob"`
	AvatarURL      string `json:"avatar_url,omitempty"`
	About          string `json:"about,omitempty"`
	ProfileSetting string `json:"profile_setting,omitempty"`
}

type AuthResponse struct {
	IsAuthenticated bool `json:"is_authenticated"`
}

type Session struct {
	Id           int       `json:"id"`
	SessionToken string    `json:"session_token"`
	UserID       int       `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Post struct {
	Id             int       `json:"id"`
	UserID         int       `json:"user_id"`
	GroupID        int       `json:"group_id,omitempty"`
	Title          string    `json:"title"`
	Content        string    `json:"content,omitempty"`
	ImageURL       string    `json:"image_url,omitempty"`
	PrivacySetting string    `json:"privacy_setting"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PostsResponse struct {
	Id             int       `json:"id"`
	UserID         int       `json:"user_id"`
	GroupID        int       `json:"group_id,omitempty"`
	Title          string    `json:"title"`
	Content        string    `json:"content,omitempty"`
	ImageURL       string    `json:"image_url,omitempty"`
	PrivacySetting string    `json:"privacy_setting"`
	CreatedAt      time.Time `json:"created_at"`
	Likes          int       `json:"likes"`
	Dislikes       int       `json:"dislikes"`
	Creator        string    `json:"creator"`
	CreatorAvatar  string    `json:"creator_avatar"`
}

type CommentsResponse struct {
	Id        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id,omitempty"`
	Content   string    `json:"content"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Username  string    `json:"username"`
	ImageURL  string    `json:"profile_image"`
}

type CreateCommentRequest struct {
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	ImageURL string `json:"image"`
}

type CreatePostRequest struct {
	PostID         int    `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"content,omitempty"`
	Username       string `json:"username"`
	GroupID        int    `json:"group_id,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
	PrivacySetting string `json:"privacy_setting"`
	CreatedAt      string `json:"created_at"`
}

type UpdatePostRequest struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"content,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
	PrivacySetting string `json:"privacy_setting"`
}

type Comment struct {
	Id        int            `json:"id,omitempty"`
	PostID    int            `json:"post_id"`
	UserID    int            `json:"user_id,omitempty"`
	Content   string         `json:"content"`
	Image     sql.NullString `json:"image,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
}

type UpdateCommentRequest struct {
	Id        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Group struct {
	Id            int           `json:"id"`
	CreatorId     int           `json:"creator_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Image         string        `json:"image"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty"`
	Members       []GroupMember `json:"members,omitempty"`
	IsUserCreator bool          `json:"is_user_creator,omitempty"`
	IsUserMember  bool          `json:"is_user_member,omitempty"`
}

type GroupMember struct {
	GroupID  int       `json:"group_id"`
	UserID   int       `json:"user_id"`
	Username string    `json:"username"`
	ImageURL string    `json:"image"`
	Status   string    `json:"status"`
	JoinedAt time.Time `json:"joined_at"`
}

type Friend struct {
	Id           int       `json:"id"`
	UserId1      int       `json:"user_id_1"`
	UserId2      int       `json:"user_id_2"`
	Status       string    `json:"status"`
	ActionUserId int       `json:"action_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FriendRequest struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// Required information for the friends list
type FriendList struct {
	UserID    int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

type Notification struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	GroupId   int       `json:"group_id,omitempty"`
	SenderId  int       `json:"sender_id,omitempty"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type GroupInvitation struct {
	Id           int       `json:"id"`
	GroupId      int       `json:"group_id"`
	JoinUserId   int       `json:"join_user_id"`
	Username     string    `json:"username"`
	ImageURL     string    `json:"image"`
	InviteUserId int       `json:"invite_user_id,omitempty"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type GroupInvitationRequest struct {
	GroupId    int `json:"group_id"`
	JoinUserId int `json:"join_user_id"`
}

type Event struct {
	Id          int       `json:"id"`
	CreatorId   int       `json:"creator_id"`
	GroupId     int       `json:"group_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}

type EventAttendance struct {
	Id        int       `json:"id"`
	EventId   int       `json:"event_id"`
	UserId    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type VoteData struct {
	Item   string `json:"item"`    // 'comment' or 'post'
	ItemID int    `json:"item_id"` // comment or post id
	Action string `json:"action"`  // 'like' or 'dislike'
}
