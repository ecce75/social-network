# IrieSphere backend

brrt

## Guide to navigating backend structure

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

------
## Functionalities

### Session

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

The endpoint will decode the data, get the user by email or username, compare the input password and stored hashed password, generate a new session token, store the session, set the sessiontoken cookie and return a success response.

-----

```go
mux.HandleFunc("/api/users/check-auth", handler.CheckAuth)
```

This endpoint will perform an auth check of the user and return a boolean value.

-----
### Posts

```go
mux.HandleFunc("/post", handler.CreatePostHandler).Methods("POST")
```

This endpoint requires post title, content, imageurl(may be empty) and privacy
setting('public', 'private', 'custom').

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

-----

### Comments

```go
mux.HandleFunc("/post/{id}/comments", handler.GetCommentsByUserIDorPostID).Methods("GET")
```

This endpoint retrieves all comments for a post by its ID. It requires the ID as a URL parameter.

-----

```go
mux.HandleFunc("/comment", handler.CreateCommentHandler).Methods("POST")
```

This endpoint creates a new comment. It requires the comment data in the request body. The user authentication is double checked via cookie and userID attached to the create comment request. After request data is decoded and stored it will return the id of the comment.

-----

```go
mux.HandleFunc("/comment/{id}", handler.DeleteCommentHandler).Methods("DELETE")
```

This endpoint deletes a comment by its ID. It requires the ID as a URL parameter. The user authentication is double checked via cookie and userID attached to the delete comment request. If the user is authorized, the comment will be deleted.
