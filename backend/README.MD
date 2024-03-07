# IrieSphere backend

## Table of Contents

- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
  - [Session](#session)
  - [Posts](#posts)
  - [Comments](#comments)
  - [Groups](#groups)
  - [Friends](#friends)
  - [Profile](#profile)
  - [Events](#events)
  - [Notifications](#notifications)
  - [Votes](#votes)
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
│   ├── model
│   │   └── # Data structures(structs)
│   └── repository
│       └── # Data access layer
├── README.md
├── server.go # Application entrypoint
└── hooks
    └── hooks.go # Utility functions that don't naturally fit elsewhere
```

## API Endpoints

### Session

- **User Registration**: Endpoint `/api/users/register` (POST)
- **User Logout**: Endpoint `/api/users/logout` (POST)
- **User Login**: Endpoint `/api/users/login` (POST)
- **Check User Authentication**: Endpoint `/api/users/check-auth` (GET)
- **Get users list**: Endpoint `/api/users/list` (GET)

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

```go
mux.HandleFunc("/api/users/list", userHandler.ListUsersHandler).Methods("GET")
```

This endpoint retrieves a list of users excluding the authenticated user.

Returns following values of users:

```go
type UserList struct {
  Id        int    `json:"id"`
  Username  string `json:"username"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  AvatarURL string `json:"avatar_url"`
}
```

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

---
---

### Groups

The Groups functionality allows for the creation, management, and deletion of user groups. It supports operations such as creating a new group, editing group details, deleting a group, and managing group memberships and invitations.

#### Group Endpoints

- **Get All Groups:** (GET) `/groups` - Retrieves all groups.
- **Create Group:** (POST) `/groups` - Allows authenticated users to create a new group.
- **Get Group by ID:** (GET) `/groups/{id}` - Retrieves details of a specific group by its ID.
- **Edit Group:** (PUT) `/groups/{id}` - Allows editing of group details.
- **Delete Group:** (DELETE) `/groups/{id}` - Deletes a group by its ID.

---

```go
mux.HandleFunc("/groups", groupHandler.GetAllGroupsHandler).Methods("GET")
```

This endpoint retrieves all groups.

---

```go
mux.HandleFunc("/groups", groupHandler.CreateGroupHandler).Methods("POST")
```

This endpoint creates a new group. Requires Title and Description in the request body. Checks for user authorization via SessionToken and attaches userID as CreatorId.

---

```go
mux.HandleFunc("/groups/{id}", groupHandler.GetGroupByIDHandler).Methods("GET")
```

This endpoint retrieves a group by its ID. It requires the ID as a URL parameter.

---

```go
mux.HandleFunc("/groups/{id}", groupHandler.EditGroupHandler).Methods("PUT")
```

This endpoint allows a user to edit a certain group. Requires the ID as URL parameter. Checks if the user is authorized to edit via SessionToken.

---

```go
mux.HandleFunc("/groups/{id}", groupHandler.DeleteGroupHandler).Methods("DELETE")
```

This endpoint allows a user to delete a certain group. Requires the ID as URL parameter. Checks if the user is authorized to delete via SessionToken.

---
---

#### Group Membership and Invitations

- **Get All Group Invitations:** (GET) `/invitations` - Gets all group invitations.
- **Invite Group Member:** (POST) `/invitations` - Enables group admins to send invitations to users.
- **Get Group Invitation By ID:** (GET) `/invitations/{id}` - Gets group invitation by ID.
- **Accept Group Membership:** (PUT) `/invitations/approve/{id}` - Allows group admins to approve membership requests.
- **Decline Group Membership:** (PUT) `/invitations/{id}` - Used by admins to decline membership requests or by users to decline invitations.
- **Request Group Membership:** (POST) `/invitations/request/{id}` - Allows users to request membership to a group.
- **Accept Group Invitation:** (PUT) `/invitations/{id}` - Allows users to accept group invitations.
- **Remove Group Member:** (DELETE) `/groups/{groupId}/members/{userId}` - Admins can remove users from their group.

---

```go
mux.HandleFunc("/invitations", groupMemberHandler.GetAllGroupInvitationsHandler).Methods("GET")
```

This endpoint retrieves all pending group invitations. Extracts userID from the cookie.

---

```go
mux.HandleFunc("/invitations", groupMemberHandler.InviteGroupMemberHandler).Methods("POST")
```

This endpoint enables users to be invited into a group. Requires groupID, inviteUserID.

The request then is processed and user authentication is double checked via cookie and userID attached to the invitation request as inviteUserID. After request data is decoded and stored it will return the id of the invitation.

---

```go
mux.HandleFunc("/invitations/{id}", groupMemberHandler.GetGroupInvitationByIDHandler).Methods("GET")

```

This endpoint provides group invitations via ID. Requires ID as URL parameter. Checks if the user is authorized via Sessiontoken.

---

```go
mux.HandleFunc("/invitations/{id}", groupMemberHandler.DeclineGroupInvitationHandler).Methods("PUT")
```

This endpoint enables users to decline a group invitation. It notifices the sender about the decline and then deletes the invitation from the database.

---

```go
mux.HandleFunc("/invitations/{id}", groupMemberHandler.AcceptGroupInvitationHandler).Methods("PUT")
```

Requires ID as URL parameter. This endpoint enables users to accept a group invitation. It notifices the sender about the decline and then deletes the invitation from the database.

---

```go
mux.HandleFunc("/invitations/request/{id}", groupMemberHandler.RequestGroupMembershipHandler).Methods("POST")
```

This endpoint enables a user to request membership to a group. Requires groupID and userID

The request then is processed and user authentication is double checked via cookie and userID attached to the request as joinUserID. After request data is decoded and stored it will notify the group admin and user of the sent request.
---

```go
mux.HandleFunc("/groups/{groupId}/members/{userId}", groupMemberHandler.RemoveMemberHandler).Methods("DELETE")
```

This endpoint enables removing a member from a group. Needs the groupID, userID as URL parameters.

Checks if there is a groupID and userID in the first place, then checks for authorization via SessionToken. Then checks if the user trying to remove another user is the group owner.

---

```go
mux.HandleFunc("/invitations/approve/{id}", groupMemberHandler.ApproveGroupMembershipHandler).Methods("PUT")
```

This endpoint enables group owner to approve a user's membership request. Expects group invitation ID as URL parameter.

Accepts the group invitation and updates the membership status to "approved."
Adds the user to the specified group.
Notifies the user of the approval.
Marks the group invitation as accepted in the database.

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
	Deleted     bool      `json:"deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

### Friends

The Friends functionality allows for the management of friend relationships between users. It supports operations such as sending friend requests, accepting friend requests, declining friend requests, blocking and unblocking users, and retrieving a user's friends.

**Send Friend Request:** (POST) `/friends/{id}` - Allows authenticated users to send a friend request to another user.
**Accept Friend Request:** (PUT) `/friends/{id}/accept` - Allows users to accept a friend request from another user.
**Decline Friend Request:** (PUT) `/friends/{id}/decline` - Allows users to decline a friend request from another user.
**Block User:** (PUT) `/friends/{id}/block` - Allows users to block another user.
**Unblock User:** (PUT) `/friends/{id}/unblock` - Allows users to unblock a user they have previously blocked.
**Get Friends:** (GET) `/friends` - Retrieves all friends of the authenticated user.

The friend request works by friends table field status(accepted, pending, declined, blocked). The handler layer uses the status field to handle different actions based on the status of the friend request.

---

```go
mux.HandleFunc("/friends/{id}", handler.SendFriendRequestHandler).Methods("POST")
```

This endpoint allows an authenticated user to send a friend request to another user. It requires the ID of the user to send the request to as a URL parameter.

---

```go
mux.HandleFunc("/friends/{id}/accept", handler.AcceptFriendRequestHandler).Methods("PUT")
```

This endpoint allows a user to accept a friend request from another user. It requires the ID of the user who sent the request as a URL parameter.

---

```go
mux.HandleFunc("/friends/{id}/decline", handler.DeclineFriendRequestHandler).Methods("PUT")
```

This endpoint allows a user to decline a friend request from another user. It requires the ID of the user who sent the request as a URL parameter.

---

```go
mux.HandleFunc("/friends/{id}/block", handler.BlockUserHandler).Methods("PUT")
```

This endpoint allows a user to block another user. It requires the ID of the user to block as a URL parameter.

---

```go
mux.HandleFunc("/friends/{id}/unblock", handler.UnblockUserHandler).Methods("PUT")
```

This endpoint allows a user to unblock a user they have previously blocked. It requires the ID of the user to unblock as a URL parameter.

---

```go
mux.HandleFunc("/friends", handler.GetFriendsHandler).Methods("GET")
```

This endpoint retrieves all friends of the authenticated user.

---

#### Friends related code

```go
type FriendList struct {
    UserID     int    `json:"user_id"`
    FirstName  string `json:"first_name"`
    LastName   string `json:"last_name"`
    AvatarURL  string `json:"avatar_url"`
    Username   string `json:"username"`
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
```

```sql
CREATE TABLE IF NOT EXISTS friends (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id1 INTEGER NOT NULL,
    user_id2 INTEGER NOT NULL,
    status TEXT CHECK(status IN ('pending', 'accepted', 'declined', 'blocked')),
    action_user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id1) REFERENCES users(id),
    FOREIGN KEY (user_id2) REFERENCES users(id),
    FOREIGN KEY (action_user_id) REFERENCES users(id)
);
```

---

### Profile

The Profile functionality allows for the retrieval of user profile information. It supports operations such as getting a user's profile and getting all posts by a user.

- **Get User Profile:** (GET) `/profile/users/{id}` - Retrieves the profile of a user by their ID.
- **Get All User Posts:** (GET) `/profile/posts/{id}` - Retrieves all posts made by a user by their ID.
- **Edit User Profile:** (PUT) `/profile/users/{id}` - Update the authenticated user profile.

---

```go
mux.HandleFunc("/profile/users/{id}", userHandler.GetUserProfileByIDHandler).Methods("GET")
```

This endpoint retrieves the profile of a user by their ID. It requires the ID of the user as a URL parameter.

---

```go
mux.HandleFunc("/profile/posts/{id}", postHandler.GetAllUserPostsHandler).Methods("GET")
```

This endpoint retrieves all posts made by a user by their ID. It requires the ID of the user as a URL parameter.
It returns users public posts and private posts if the requesting user is friends with the target user.
Doesn't retrieve group posts.

---

```go
mux.HandleFunc("/profile/users/{id}", userHandler.EditUserProfileHandler).Methods("PUT")
```

This endpoint updates the authenticated user profile by userid from cookie. Introduced ProfileSetting field to RegistrationData struct.

```go
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
```

---

#### Profile related code

```go
type Profile struct {
    Id        int    `json:"id"`
    Username  string `json:"username"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    DOB       string `json:"dob"`
    AvatarURL string `json:"avatar_url"`
    About     string `json:"about"`
    CreatedAt string `json:"created_at"`
}
```

User handlers formvalues:

```go
regData.Username = r.FormValue("username")
regData.Email = r.FormValue("email")
regData.Password = r.FormValue("password")
regData.FirstName = r.FormValue("first_name")
regData.LastName = r.FormValue("last_name")
regData.DOB = r.FormValue("dob")
regData.About = r.FormValue("about")
regData.ProfileSetting = r.FormValue("profile_setting")
```

User table:

```sql
profile TEXT NOT NULL CHECK(profile IN ('public', 'private')) DEFAULT 'public',
```

---

---

### Events

The Event functionality encompasses operations related to events, offering functionalities such as creating, retrieving, updating, and deleting events. Additionally, it supports managing attendance records for events.

---

- **Get All Events:** (GET) `/events` - Retrieves a list of all events.
- **Create Event:** (POST) `/events` - Creates a new event. Requires authentication.
- **Get Event by ID:** (GET) `/events/{id}` - Retrieves details of a specific event by its ID.
- **Edit Event:** (PUT) `/events/{id}` - Updates details of a specific event. Requires authentication and ownership verification.
- **Delete Event:** (DELETE) `/events/{id}` - Deletes a specific event. Requires authentication and ownership verification.
- **Get Events by Group ID:** (GET) `/events/{id}` - Retrieves events associated with a specific group.
- **Add or Update Attendance:** (PUT) `/events/{id}` - Adds or updates attendance for a user at a specific event. Requires authentication.
- **Get Attendance by Event ID:** (GET) `/events/{id}` - Retrieves attendance records for a specific event.

---

```go
mux.HandleFunc("/events", eventHandler.GetAllEventsHandler).Methods("GET")
```

This endpoint retrieves all events. 

---

```go
mux.HandleFunc("/events", eventHandler.CreateEventHandler).Methods("POST")
```

This endpoint processes an HTTP request to create a new event. It decodes the request body, authenticates the user through the SessionToken, checks group membership, sets the creator ID, and stores the event in the database. The response includes a status code indicating success or specific error details if applicable.

---

```go
mux.HandleFunc("/events/{id}", eventHandler.GetEventByIDHandler).Methods("GET")
```

This endpoint processes an HTTP request to retrieve details for a specific event based on its ID. It parses the event ID from the request parameters, validates it, retrieves the event details from the repository, and responds with the event data in JSON format. Error responses cover issues such as invalid event IDs or errors in retrieving event details, ensuring appropriate handling of potential failures.

---

```go
mux.HandleFunc("/events/{id}", eventHandler.EditEventHandler).Methods("PUT")
```

This endpoint enables a user to edit an existing event. Decodes the request body, confirms user authentication via ID from request parameters, checks if the user is the event creator, updates the event in the database, and responds with the updated event details in JSON format upon success. Error responses are provided for decoding issues, authentication failures, unauthorized access, or database update errors.

---

```go
mux.HandleFunc("/events/{id}", eventHandler.DeleteEventHandler).Methods("DELETE")
```

This endpoint processes an HTTP request to delete an event. It confirms user authentication, retrieves and validates the event ID from request parameters, checks if the user is the event creator, deletes the event from the database, and responds with a success message in JSON format. Error responses cover issues such as authentication errors, invalid event IDs, unauthorized access, or database deletion failures.

---

```go
mux.HandleFunc("/events/{id}", eventHandler.GetEventsByGroupIDHandler).Methods("GET")
```

This endpoint handles an HTTP request to retrieve events based on a group ID. It parses the group ID from the request parameters, validates it, retrieves the associated events from the repository, and responds with the events in JSON format. Error responses cover issues such as parsing errors, missing group ID parameters, or failures in retrieving events from the repository.

---

```go
mux.HandleFunc("/events/{id}", eventHandler.AddOrUpdateAttendanceHandler).Methods("PUT")
```

This endpoint processes an HTTP request to add or update attendance for a specific event and user. It extracts event ID, user ID, and attendance status from the request parameters, validates them, and invokes the corresponding repository method. The function responds with the attendance ID in JSON format upon success, handling errors such as missing parameters or failures in parsing and updating attendance with appropriate error responses.

---

```go
mux.HandleFunc("/events/{id}", eventHandler.GetAttendanceByEventIDHandler).Methods("GET")
```

This endpoint handles an HTTP request to retrieve attendance for a specific event. It parses the event ID from the request parameters, validates it, retrieves the attendance list from the repository, and responds with the attendance data in JSON format. Error responses cover issues such as missing event ID parameters or parsing errors, ensuring appropriate handling of potential failures.

---

#### Events related code

```go
type Event struct {
  Id          int       `json:"id"`
  CreatorId   int       `json:"creator_id"`
  Title       string    `json:"title"`
  Description string    `json:"description"`
  Location    string    `json:"location"`
  StartTime   time.Time `json:"start_time"`
  EndTime     time.Time `json:"end_time"`
  CreatedAt   time.Time `json:"created_at"`
}
```

```go
type EventAttendance struct {
  Id        int       `json:"id"`
  EventId   int       `json:"event_id"`
  UserId    int       `json:"user_id"`
  Status    string    `json:"status"`
  CreatedAt time.Time `json:"created_at"`
}
```

---

---

### Notifications

The Notification functionality facilitates operations related to user notifications, including retrieving, creating, and managing notification status.

---

- **Get All Notifications:** (GET) /notifications - Retrieves a list of all notifications for the authenticated user.
- **Create Notification:** (POST) /notifications - Creates a new notification. Requires authentication.
- **Get Notification by ID:** (GET) /notifications/{id} - Retrieves details of a specific notification by its ID.
- **Mark Notification as Read:** (PUT) /notifications/{id} - Marks a specific notification as read. Requires authentication.

---

```go
mux.HandleFunc("/notifications", notificationHandler.GetAllNotificationsHandler).Methods("GET")
```

This endpoint manages an HTTP request to retrieve all notifications. It fetches the notifications from the repository, responds with the data in JSON format, and handles potential errors, providing appropriate error responses for failures in retrieving notifications.

---

```go
mux.HandleFunc("/notifications", notificationHandler.CreateNotificationHandler).Methods("POST")
```

This endpoint handles an HTTP request to create a new notification. It decodes the request body, validates user authentication, assigns the user ID, and creates the notification in the database. The function responds with a success status (201 Created) or appropriate error messages for issues such as decoding errors or database creation failures.

---

```go
mux.HandleFunc("/notifications/{id}", notificationHandler.GetNotificationByIDHandler).Methods("GET")
```

This endpoint processes an HTTP request to retrieve details for a specific notification based on its ID. It parses the notification ID from the request parameters, validates it, retrieves the notification details from the repository, and responds with the data in JSON format. Error responses cover issues such as invalid notification IDs or errors in retrieving notification details, ensuring appropriate handling of potential failures.

---

```go
mux.HandleFunc("/notifications/{id}", notificationHandler.MarkNotificationAsReadHandler).Methods("PUT")
```

This endpoint handles an HTTP request to mark a notification as read. It parses the notification ID from the request parameters, validates it, and updates the notification status in the repository. The function responds with a success status (200 OK) or appropriate error messages for issues such as invalid notification IDs or failures in updating the notification status.

---

#### Notifications related code

```go
type Notification struct {
  Id        int       `json:"id"`
  UserId    int       `json:"user_id"`
  Type      string    `json:"type"`
  Message   string    `json:"message"`
  IsRead    bool      `json:"is_read"`
  CreatedAt time.Time `json:"created_at"`
}
```

```sql

```

---

### Votes

- **Vote post or comment** (POST) `/vote` - Like or dislike a comment/post, removing the vote works through same endpoint by sending the already selected vote again.

```go
mux.HandleFunc("/vote", voteHandler.VotePostOrCommentHandler).Methods("POST")
```

The endpoint requires 3 values: item, item_id, action

- item - 'comment' or 'post'
- item_id - id of the comment or post
- action - 'like' or 'dislike'

---

Logic behind votes:

- post request with all required fields
- check user auth
- check if item and action values are correct
- call the repository function
- repository level checks if the vote for specific comment or post exists
- if vote doesn't exist -> create it
- if it exists -> check if same type(action)
- if same type -> remove the vote
- if different type(change like to dislike for example) -> update the vote

---

#### Votes related code

```go
type VoteData struct {
  Item   string `json:"item"`    // 'comment' or 'post'
  ItemID int    `json:"item_id"` // comment or post id
  Action string `json:"action"`  // 'like' or 'dislike'
}
```

---

---

## Backend contribution

fork -> contribute -> pull request

Notes:

- Use repository pattern for data access layer.

- Document everything, clear commit messages about changes

- Don't forget Jira issue links(ex. SN-44) for commits and pull requests

## Future Work

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
