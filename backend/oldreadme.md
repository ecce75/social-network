# IrieSphere backend

brrt

## Table of Contents

- [Backend Structure](#backend-structure)
  - [pkg](#pkg)
    - [db](#db)
    - [model](#model)
    - [repository](#repository)
    - [handler](#handler)
  - [api](#api)
  - [util](#util)
- [Functionalities](#functionalities)
  - [Session](#session)
  - [Posts](#posts)
  - [Comments](#comments)
  - [Groups](#groups)
    - [Groups](#groups-todo)

## Backend structure

### pkg

This directory contains the core logic of your application and is where most of your code will reside.

```bash
./backend
├── pkg
│   │   ├── db
│   │   │   ├── database.db # Database file, don't touch if not sure what you're doing!
│   │   │   ├── migrations
│   │   │   │   └── sqlite # SQL queryset
│   │   │   │       ├── 000001_create_users_table.down.sql
│   │   │   │       └── 000001_create_users_table.up.sql
│   │   │   └── sqlite
│   │   │       └── sqlite.go
│   │   ├── model
│   │   │   └── # Data structures/structs
│   │   ├── repository
│   │   │   └── # Data access layer
│   │   └── handler
│   │       └──  # Handlers call into repositories to fetch and store data
```

#### db

Contains database setup, connection logic, and migrations. Keeping migrations close to your database logic makes sense since they are tightly coupled.

#### model

Defines the data structures used by your application. These are often representations of your database tables but can also include other domain-specific logic. Structs basically...

#### repository

Acts as the data access layer. Repositories use models to interact with the database, encapsulating the logic needed to query or update the database.

#### handler

Contains the business logic of your application. Handlers call into repositories to fetch and store data, perform operations, and implement application-specific rules.

Handlers should also validate and sanitize the input data.

### api

```bash
./backend
├── api
│   └── router.go # HTTP handlers and routing
```

This is where you define your HTTP handlers and routing. It's the layer that interacts with the outside world, translating HTTP requests into actions on your services and ultimately your models and database.

### util

A place for utility functions that don't naturally fit elsewhere. These might include helper functions for formatting, validation, or small tasks used across multiple parts of the application.

- Session Token Generation

-----

## Functionalities

### Session

```go
type Session struct {
 Id     int   `json:"id"`
 SessionToken  string   `json:"session_token"`
 UserID    int   `json:"user_id"`
 ExpiresAt   time.Time  `json:"expires_at"`
}
```

- **User Registration**: Endpoint `/api/users/register` (POST)
- **User Logout**: Endpoint `/api/users/logout` (POST)
- **User Login**: Endpoint `/api/users/login` (POST)
- **Check User Authentication**: Endpoint `/api/users/check-auth` (GET)

-----

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

It will then decode the request data, hash the password, store the user in database, generate sessionToken, set the sessionToken cookie and return a success response.

-----

```go
mux.HandleFunc("/api/users/logout", handler.LogoutHandler).Methods("POST")
```

Logout gets the session token from cookie and deletes it.

-----

```go
mux.HandleFunc("/api/users/login", handler.LoginHandler).Methods("POST")
```

Using this endpoint requires:

- username (could aswell be email)
- password

```go
type LoginData struct {
 Username string `json:"username"`
 Password string `json:"password"`
}
```

The endpoint will decode the data, get the user by email or username, compare the input password and stored hashed password, generate a new session token, store the session, set the sessiontoken cookie and return a success response.

-----

```go
mux.HandleFunc("/api/users/check-auth", handler.CheckAuth)
```

```go
type AuthResponse struct {
 IsAuthenticated bool `json:"is_authenticated"`
}
```

This endpoint will perform an auth check of the user and return a boolean value.

-----

### Posts

```go
type Post struct {
 Id     int   `json:"id"`
 UserID    int   `json:"user_id"`
 Title   string   `json:"title"`
 Content   string   `json:"content,omitempty"`
 ImageURL   string   `json:"image_url,omitempty"`
 PrivacySetting  string     `json:"privacy_setting"`
    CreatedAt       time.Time  `json:"created_at"`
}
```

- **Create Post**: Endpoint `/post` (POST)
- **Delete Post**: Endpoint `/post/{id}` (DELETE)
- **Update Post**: Endpoint `/post/{id}` (PUT)

-----

```go
mux.HandleFunc("/post", handler.CreatePostHandler).Methods("POST")
```

This endpoint requires post title, content, imageurl(may be empty) and privacy
setting('public', 'private', 'custom').

```go
type CreatePostRequest struct {
 Title    string `json:"title"`
 Content   string `json:"content,omitempty"`
 ImageURL   string `json:"image_url,omitempty"`
 PrivacySetting  string `json:"privacy_setting"`
}
```

The request then is processed and user authentication is double checked via cookie and userID attached to the create post request. After request data is decoded and stored it will return the id of the post.

-----

```go
mux.HandleFunc("/post/{id}", handler.DeletePostHandler).Methods("DELETE")
```

This endpoint deletes a post by its ID. It requires the ID as a URL parameter.

-----

```go
mux.HandleFunc("/post/{id}", handler.UpdatePostHandler).Methods("PUT")
```

This endpoint updates a post by its ID. It requires the ID as a URL parameter and the new post data in the request body.

```go
type UpdatePostRequest struct {
 Id int `json:"id"`
 Title string `json:"title"`
 Content string `json:"content,omitempty"`
 ImageURL string `json:"image_url,omitempty"`
 PrivacySetting string `json:"privacy_setting"`
}
```

-----

### Comments

```go
type Comment struct {
 Id int `json:"id"`
 PostID int `json:"post_id"`
 UserID int `json:"user_id"`
 Content string `json:"content"`
 CreatedAt time.Time `json:"created_at"`
}
```

- **Get Comments**: Endpoint `/post/{id}/comments` (GET)
- **Create Comment**: Endpoint `/comment` (POST)
- **Delete Comment**: Endpoint `/comment/{id}` (DELETE)

-----

```go
mux.HandleFunc("/post/{id}/comments", handler.GetCommentsByUserIDorPostID).Methods("GET")
```

This endpoint retrieves all comments for a post by its ID. It requires the ID as a URL parameter.

-----

```go
mux.HandleFunc("/comment", handler.CreateCommentHandler).Methods("POST")
```

This endpoint creates a new comment. It requires the comment data in the request body. The user authentication is double checked via cookie and userID attached to the create comment request. After request data is decoded and stored it will return the id of the comment.

```go
type Comment struct {
 Id int `json:"id"`
 PostID int `json:"post_id"`
 UserID int `json:"user_id"`
 Content string `json:"content"`
 CreatedAt time.Time `json:"created_at"`
}
```

!!! need to create new comment struct because id won't be available before creation !!!

-----

```go
mux.HandleFunc("/comment/{id}", handler.DeleteCommentHandler).Methods("DELETE")
```

This endpoint deletes a comment by its ID. It requires the ID as a URL parameter. The user authentication is double checked via cookie and userID attached to the delete comment request. If the user is authorized, the comment will be deleted.

-----

### Groups

```go
type Group struct {
 Id     int   `json:"id"`
 Name   string   `json:"name"`
 Description   string   `json:"description,omitempty"`
 CreatedAt   time.Time  `json:"created_at"`
}
```

- **Delete Group:** Endpoint /group/{id} (DELETE)

-----

```go
mux.HandleFunc("/group/{id}", handler.DeleteGroupHandler).Methods("DELETE")
```

This endpoint deletes a group by its ID. It requires the ID as a URL parameter. When a group is deleted, all the links to its members are also deleted. This is achieved by modifying the SQL schema to delete on cascade:

```sql
CREATE TABLE group_members (
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

With this modification, the database automatically deletes all the referencing records in the group_members table when the referenced record in the groups table is deleted.

The DeleteGroupHandler also logs a message to the console whenever a group is successfully deleted. The message includes the ID of the deleted group.

-----

- **Get Group by ID**: Endpoint `/group/{id}` (GET)

```go
mux.HandleFunc("/group/{id}", handler.GetGroupHandler).Methods("GET")
```

This endpoint retrieves a group by its ID. It requires the ID as a URL parameter. If no group with the given ID exists, it returns a 404 Not Found error.

-----

- **Update Group:** Endpoint /group/{id} (PUT)

```go
mux.HandleFunc("/group/{id}", handler.UpdateGroupHandler).Methods("PUT")
```

This endpoint updates a group's details. It requires the ID as a URL parameter and the new details as a JSON body. If no group with the given ID exists, it returns a 404 Not Found error.

-----

#### Groups todo

- Implement: group invitation, group members crud, logging, notifications...
- Test: endpoint tests, sql test data.
