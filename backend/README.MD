# IrieSphere backend

## Table of Contents

- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
  - [Session](#session)
  - [Posts](#posts)
  - [Comments](#comments)
  - [Groups](#groups)
- [Backend contribution](#backend-contribution)
- [Future Work](#future-work)
- [Extra](#extra)
  - [Migrate](#migrate)
  - [Repository pattern](#repository-pattern)

## Project Structure

The backend is structured into several packages, each with a specific purpose:

- `pkg`: Contains the core logic of the application.
  - `db`: Handles database setup, connection logic, and migrations.
  - `model`: Defines the data structures used by the application.
  - `repository`: Acts as the data access layer, using models to interact with the database.
  - `handler`: Contains the business logic of the application, calling into repositories to fetch and store data.
- `api`: Defines HTTP handlers and routing.
- `util`: Contains utility functions used across the application.

```bash
.
├── api
│   └── router.go # Routes/API Endpoints
├── pkg
│   ├── db
│   │   ├── database.db # Database file
│   │   ├── migrations
│   │   │   └── sqlite # SQL Migrations query set
│   │   │       ├── 000001_create_users_table.down.sql
│   │   │       └── 000001_create_users_table.up.sql
│   │   └── sqlite
│   │       └── sqlite.go
│   ├── handler
│   │   └── # Handlers do the magic of organizing everything
│   ├── middleware
│   │   └── # Middleware for apis
│   ├── model
│   │   └── # Data structures(structs)
│   └── repository
│       └── # Data access layer
├── README.md
├── server.go # Application entrypoint
└── util
    └── util.go # Utility functions that don't naturally fit elsewhere
```

## API Endpoints

### Session

- **User Registration**: Endpoint `/api/users/register` (POST)
- **User Logout**: Endpoint `/api/users/logout` (POST)
- **User Login**: Endpoint `/api/users/login` (POST)
- **Check User Authentication**: Endpoint `/api/users/check-auth` (GET)

---

```go
mux.HandleFunc("/api/users/register", handler.UserRegisterHandler).Methods("POST")
```

Using this endpoint requires:

- username
- email
- password
- first_name
- last_name
- dob (date of birth)
- avatar_url (omitempty)
- about

It will then decode the request data, hash the password, store the user in database, generate sessionToken, set the sessionToken cookie and return a success response.

---

```go
mux.HandleFunc("/api/users/logout", handler.LogoutHandler).Methods("POST")
```

Logout gets the session token from cookie and deletes it.

---

```go
mux.HandleFunc("/api/users/login", handler.LoginHandler).Methods("POST")
```

Using this endpoint requires:

- username (could aswell be email)
- password

The endpoint will decode the data, get the user by email or username, compare the input password and stored hashed password, generate a new session token, store the session, set the sessiontoken cookie and return a success response.

---

```go
mux.HandleFunc("/api/users/check-auth", handler.CheckAuth)
```

This endpoint will perform an auth check of the user and return a boolean value.

---

#### Session related code

```go
type Session struct {
 Id     int   `json:"id"`
 SessionToken  string   `json:"session_token"`
 UserID    int   `json:"user_id"`
 ExpiresAt   time.Time  `json:"expires_at"`
}
```

```go
type RegistrationData struct {
 Username  string `json:"username"`
 Email   string `json:"email"`
 Password  string `json:"password"`
 FirstName  string `json:"first_name"`
 LastName  string `json:"last_name"`
 DOB   string `json:"dob"`
 AvatarURL  string `json:"avatar_url,omitempty"`
 About   string `json:"about,omitempty"`
}
```

```go
type LoginData struct {
 Username string `json:"username"`
 Password string `json:"password"`
}
```

```go
type AuthResponse struct {
 IsAuthenticated bool `json:"is_authenticated"`
}
```

### Posts

- **Get All Posts**: Endpoint `/post` (GET)
- **Get Posts By Group ID**: Endpoint `/groups/posts/{id}` (GET)
- **Create Post**: Endpoint `/post` (POST)
- **Delete Post**: Endpoint `/post/{id}` (DELETE)
- **Update Post**: Endpoint `/post/{id}` (PUT)

---

```go
mux.HandleFunc("/post", handler.CreatePostHandler).Methods("POST")
```

This endpoint requires post title, content, imageurl(may be empty) and privacy
setting('public', 'private', 'custom').

The request then is processed and user authentication is double checked via cookie and userID attached to the create post request. After request data is decoded and stored it will return the id of the post.

---

```go
mux.HandleFunc("/post", postHandler.GetAllPostsHandler).Methods("GET")
```

This endpoint retrieves all posts that the authenticated user has access to. It includes all public posts and posts from the user's groups.

---

```go
mux.HandleFunc("/groups/posts/{id}", postHandler.GetPostsByGroupIDHandler).Methods("GET")
```

This endpoint retrieves all posts for a specific group by its ID. It requires the ID as a URL parameter.

---

```go
mux.HandleFunc("/post/{id}", handler.DeletePostHandler).Methods("DELETE")
```

This endpoint deletes a post by its ID. It requires the ID as a URL parameter.

---

```go
mux.HandleFunc("/post/{id}", handler.UpdatePostHandler).Methods("PUT")
```

This endpoint updates a post by its ID. It requires the ID as a URL parameter and the new post data in the request body.

---

#### Post related code

```go
type Post struct {
 Id     int   `json:"id"`
 UserID    int   `json:"user_id"`
 GroupID int `json:"group_id"`
 Title   string   `json:"title"`
 Content   string   `json:"content,omitempty"`
 ImageURL   string   `json:"image_url,omitempty"`
 PrivacySetting  string     `json:"privacy_setting"`
 CreatedAt       time.Time  `json:"created_at"`
}
```

```go
type CreatePostRequest struct {
 GroupID int `json:"group_id"`
 Title    string `json:"title"`
 Content   string `json:"content,omitempty"`
 ImageURL   string `json:"image_url,omitempty"`
 PrivacySetting  string `json:"privacy_setting"`
}
```

```go
type UpdatePostRequest struct {
 Id int `json:"id"`
 Title string `json:"title"`
 Content string `json:"content,omitempty"`
 ImageURL string `json:"image_url,omitempty"`
 PrivacySetting string `json:"privacy_setting"`
}
```

### Comments

- **Get Comments**: Endpoint `/post/{id}/comments` (GET)
- **Create Comment**: Endpoint `/comment` (POST)
- **Delete Comment**: Endpoint `/comment/{id}` (DELETE)

---

```go
mux.HandleFunc("/post/{id}/comments", handler.GetCommentsByUserIDorPostID).Methods("GET")
```

This endpoint retrieves all comments for a post by its ID. It requires the ID as a URL parameter.

---

```go
mux.HandleFunc("/comment", handler.CreateCommentHandler).Methods("POST")
```

This endpoint creates a new comment. It requires the comment data in the request body. The user authentication is double checked via cookie and userID attached to the create comment request. After request data is decoded and stored it will return the id of the comment.

---

#### Comments related code

```go
type Comment struct {
 Id int `json:"id"`
 PostID int `json:"post_id"`
 UserID int `json:"user_id"`
 Content string `json:"content"`
 CreatedAt time.Time `json:"created_at"`
}
```

### Groups

The Groups functionality allows for the creation, management, and deletion of user groups. It supports operations such as creating a new group, editing group details, deleting a group, and managing group memberships and invitations.

#### Group Endpoints

- **Create Group:** POST /groups - Allows authenticated users to create a new group.
- **Get All Groups:** GET /groups - Retrieves all groups.
- **Get Group by ID:** GET /groups/{id} - Retrieves details of a specific group by its ID.
- **Edit Group:** PUT /groups/{id} - Allows editing of group details.
- **Delete Group:** DELETE /groups/{id} - Deletes a group by its ID.

---

---

#### Group Membership and Invitations

- **Invite Group Member:** POST /invitations - Enables group admins to send invitations to users.
- **Approve Group Membership:** PUT /invitations/approve/{id} - Allows group admins to approve membership requests.
- **Decline Group Membership:** PUT /invitations/{id} - Used by admins to decline membership requests or by users to decline invitations.
- **Request Group Membership:** POST /invitations/request/{id} - Allows users to request membership to a group.
- **Accept Group Invitation:** PUT /invitations/{id} - Allows users to accept group invitations.
- **Remove Group Member:** DELETE /groups/{groupId}/members/{userId} - Admins can remove users from their group.

---

---

#### Groups related code

```go
type GroupInvitation struct {
 Id           int       `json:"id"`
 GroupId      int       `json:"group_id"`
 JoinUserId   int       `json:"join_user_id"`
 InviteUserId int       `json:"invite_user_id,omitempty"`
 Status       string    `json:"status"`
 CreatedAt    time.Time `json:"created_at"`
}
```

```go
type Group struct {
 Id          int       `json:"id"`
 CreatorId   int       `json:"creator_id"`
 Title       string    `json:"title"`
 Description string    `json:"description"`
 CreatedAt   time.Time `json:"created_at"`
}
```

```go
type GroupMember struct {
 GroupId  int       `json:"group_id"`
 UserId   int       `json:"user_id"`
 JoinedAt time.Time `json:"joined_at"`
}
```

---

## Backend contribution

fork -> contribute -> pull request

Notes:

- Use repository pattern for data access layer.

- Use auth middleware for requests that should require authorization.

- Document everything, clear commit messages about changes

- Don't forget Jira issue links(ex. SN-44) for commits and pull requests

## Future Work

- refactor
- group posts & comments
- events + going/not going
- friends + requests
- sanitize input data
- notification system
- chat
- tests?
- dockerize

## Extra

### Migrate

The migrate tool is used for managing database migrations. Migrations are stored in the pkg/db/migrations directory and can be applied or rolled back using the migrate CLI or programmatically within the Go application.

Useful links:

- <https://github.com/golang-migrate/migrate>

### Repository pattern

The repository pattern is used to abstract the data layer in the application, providing a cleaner way to access data. It involves using repository interfaces/types and concrete implementations that interact with the database, allowing for easier testing and maintenance.

Useful links:

- <https://github.com/jorzel/go-repository-pattern>
