package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	postRepo        *repository.PostRepository
	sessionRepo     *repository.SessionRepository
	friendsRepo     *repository.FriendsRepository
	groupMemberRepo *repository.GroupMemberRepository
	userRepo 	  	*repository.UserRepository
	voteHandler     *VoteHandler
}

func NewPostHandler(postRepo *repository.PostRepository, sessionRepo *repository.SessionRepository, friendsRepo *repository.FriendsRepository, groupMemberRepo *repository.GroupMemberRepository, userRepo *repository.UserRepository, voteHandler *VoteHandler) *PostHandler {
	return &PostHandler{postRepo: postRepo, sessionRepo: sessionRepo, friendsRepo: friendsRepo, groupMemberRepo: groupMemberRepo, userRepo: userRepo, voteHandler: voteHandler}
}

func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseMultipartForm(10 << 20) // Maximum memory 10MB, change this based on your requirements
	if err1 != nil {
		http.Error(w, "Error parsing form data: "+err1.Error(), http.StatusBadRequest)
		return
	}
	var request model.CreatePostRequest
	request.Title = r.FormValue("title")
	request.Content = r.FormValue("content")
	request.GroupID, _ = strconv.Atoi(r.FormValue("group"))
	request.PrivacySetting = r.FormValue("privacy-setting")

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming authentication: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Creates the post in database
	post, err := h.postRepo.CreatePost(&request, userID)
	if err != nil {
		http.Error(w, "Failed to create the post: "+err.Error(), http.StatusInternalServerError)
		return
	}
	util.ImageSave(w, r, strconv.Itoa(post.PostID), "post")
	// Successful response
	response := map[string]interface{}{
		"message": "Post created successfully",
		"data":    post,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) EditPostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body for updating the post
	var request model.UpdatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to decode request data", http.StatusBadRequest)
		return
	}

	// Confirm user auth and get userid
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Update the post in the database
	err = h.postRepo.UpdatePost(request.Id, userID, request)
	if err != nil {
		http.Error(w, "Failed to update the post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]string{
		"message": "Post updated successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the post ID from the URL
	vars := mux.Vars(r)
	postID, ok := vars["id"]
	intpostID, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "Failed to parse post ID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Post ID is missing in parameters", http.StatusBadRequest)
		return
	}

	userId, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Delete the post from the database
	err = h.postRepo.DeletePost(intpostID, userId)
	if err != nil {
		http.Error(w, "Failed to delete the post: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Successful response
	response := map[string]string{
		"message": "Post deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PostHandler) GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// userGroupsPosts, err := h.postRepo.GetPostsByUserGroups(userID)
	// if err != nil {
	// 	http.Error(w, "Failed to retrieve posts by user groups: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	posts, err := h.postRepo.GetAllPostsWithUserIDAccess(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve posts: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// posts = append(posts, userGroupsPosts...)

	// Append the votes to the posts
	postsResponse, err := h.voteHandler.AppendVotesToPostsResponse(posts)
	if err != nil {
		http.Error(w, "Failed to append votes to posts: "+err.Error(), http.StatusInternalServerError)
		return
	}
	for i, post := range postsResponse {
		creatorId, err := h.postRepo.GetPostOwnerIDByPostID(post.Id); if err != nil {
			http.Error(w, "Failed to retrieve post owner: "+err.Error(), http.StatusInternalServerError)
			return
		}
		creatorProfile, err := h.userRepo.GetUserProfileByID(int(creatorId))
		if err != nil {
			http.Error(w, "Failed to retrieve post owner profile: "+err.Error(), http.StatusInternalServerError)
			return
		}
		postsResponse[i].Creator = creatorProfile.Username
		postsResponse[i].CreatorAvatar = creatorProfile.AvatarURL
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postsResponse)
}

// GetAllUserPosts retrieves all posts for a specific user.
// It takes the user ID from the request parameters and checks the user's authentication.
// If the requesting user is the same as the user ID in the parameters, it retrieves all posts for that user.
// If the requesting user is not the same, it checks the friend status between the requesting user and the user ID in the parameters.
// If they are friends, it retrieves all posts for that user.
// If they are not friends, it retrieves only the public posts for that user.
// The retrieved posts are encoded as JSON and sent in the response.
func (h *PostHandler) GetAllUserPostsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, ok := vars["id"]

	requestingUserID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: "+err.Error(), http.StatusUnauthorized)
		return
	}
	var intUserID int
	if userID == "me" {
		intUserID = requestingUserID
	} else {
		intUserID, err = strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "Failed to parse user ID: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if !ok {
		http.Error(w, "User ID is missing in parameters", http.StatusBadRequest)
		return
	}

	var posts []model.Post
	if requestingUserID == intUserID {
		posts, err = h.postRepo.GetAllUserPosts(requestingUserID)
		if err != nil {
			http.Error(w, "Failed to retrieve all user posts: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// check if the users are friends
		status, err := h.friendsRepo.GetFriendStatus(requestingUserID, intUserID)
		if err != nil {
			http.Error(w, "Failed to retrieve friend status: "+err.Error(), http.StatusInternalServerError)
			return
		}
		// retrieve users public posts, and private posts if they are friends
		if status == "accepted" {
			posts, err = h.postRepo.GetAllUserPosts(intUserID)
			if err != nil {
				http.Error(w, "Failed to retrieve friend's posts: "+err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			posts, err = h.postRepo.GetAllUserPublicPosts(intUserID)
			if err != nil {
				http.Error(w, "Failed to retrieve user public posts: "+err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}

	// Append the votes to the posts
	postsResponse, err := h.voteHandler.AppendVotesToPostsResponse(posts)
	if err != nil {
		http.Error(w, "Failed to append votes to posts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i, post := range postsResponse {
		creatorId, err := h.postRepo.GetPostOwnerIDByPostID(post.Id); if err != nil {
			http.Error(w, "Failed to retrieve post owner: "+err.Error(), http.StatusInternalServerError)
			return
		}
		creatorProfile, err := h.userRepo.GetUserProfileByID(int(creatorId))
		if err != nil {
			http.Error(w, "Failed to retrieve post owner profile: "+err.Error(), http.StatusInternalServerError)
			return
		}
		postsResponse[i].Creator = creatorProfile.Username
		postsResponse[i].CreatorAvatar = creatorProfile.AvatarURL
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postsResponse)
}

// ---------------------------------------------- //
// ------------ Group Posts Handlers ------------ //
// ---------------------------------------------- //

func (h *PostHandler) GetPostsByGroupIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["groupId"])
	if err != nil {
		http.Error(w, "Group ID is missing in parameters", http.StatusBadRequest)
		return
	}
	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "Error confirming user authentication: "+err.Error(), http.StatusUnauthorized)
		return
	}
	// check if user is in group
	isMember, err := h.groupMemberRepo.IsUserGroupMember(userID, groupID)
	if err != nil {
		http.Error(w, "Failed to check if user is in group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !isMember {
		response := map[string]string{
			"message": "User not member of group",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	posts, err := h.postRepo.GetPostsByGroupID(groupID)
	if err != nil && err != sql.ErrNoRows {
		http.Error(w, "Failed to retrieve posts: "+err.Error(), http.StatusInternalServerError)
		return
	} else if err == sql.ErrNoRows {
		response := map[string]string{
			"message": "No posts found for the group",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Append the votes to the posts
	postsResponse, err := h.voteHandler.AppendVotesToPostsResponse(posts)
	if err != nil {
		http.Error(w, "Failed to append votes to posts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i, post := range postsResponse {
		creatorId, err := h.postRepo.GetPostOwnerIDByPostID(post.Id); if err != nil {
			http.Error(w, "Failed to retrieve post owner: "+err.Error(), http.StatusInternalServerError)
			return
		}
		creatorProfile, err := h.userRepo.GetUserProfileByID(int(creatorId))
		if err != nil {
			http.Error(w, "Failed to retrieve post owner profile: "+err.Error(), http.StatusInternalServerError)
			return
		}
		postsResponse[i].Creator = creatorProfile.Username
		postsResponse[i].CreatorAvatar = creatorProfile.AvatarURL
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postsResponse)
}

