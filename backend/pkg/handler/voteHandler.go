package handler

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"backend/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type VoteHandler struct {
	voteRepo *repository.VoteRepository
	sessionRepo *repository.SessionRepository
}

func NewVoteHandler(voteRepo *repository.VoteRepository, sessionRepo *repository.SessionRepository) *VoteHandler {
	return &VoteHandler{voteRepo: voteRepo, sessionRepo: sessionRepo}
}

func (h *VoteHandler) VotePostOrCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID, err := h.sessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	voteData := model.VoteData{}

	if err := json.NewDecoder(r.Body).Decode(&voteData); err != nil {
		http.Error(w, "Failed to parse request data", http.StatusBadRequest)
		return
	}
	// check input from request body
	if voteData.Item != "post" && voteData.Item != "comment" {
		http.Error(w, "Invalid item type", http.StatusBadRequest)
		return
	} else if voteData.Action != "like" && voteData.Action != "dislike" {
		http.Error(w, "Invalid action", http.StatusBadRequest)
		return
	}

	err = h.voteRepo.VoteItem(voteData, userID); if err != nil {
		errmsg := fmt.Sprintf("Failed to vote %s: %s", voteData.Item, err.Error())
		http.Error(w, errmsg, http.StatusInternalServerError)
	}
	
	var likes, dislikes, getVoteError = h.voteRepo.GetItemVotes(voteData.Item, voteData.ItemID); if getVoteError != nil {
		http.Error(w, "Failed to fetch updated votes: " + err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response struct for the votes
	response := struct {
		Likes   int `json:"likes"`
		Dislikes int `json:"dislikes"`
	}{
		Likes:   likes,
		Dislikes: dislikes,
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseData)
	if err != nil {
		// Handle error writing response here
		return
	}
}

func (h *VoteHandler) AppendVotesToPostsResponse(posts []model.Post) ([]model.PostsResponse, error) {
	postsResponse := make([]model.PostsResponse, len(posts))
	for i, post := range posts {
		likes, dislikes, err := h.voteRepo.GetItemVotes("post", post.Id)
		if err != nil {
			return nil, err
		}
		postsResponse[i] = model.PostsResponse{
            Id:             post.Id,
            UserID:         post.UserID,
            GroupID:        post.GroupID,
            Title:          post.Title,
            Content:        post.Content,
            ImageURL:       post.ImageURL,
            PrivacySetting: post.PrivacySetting,
            CreatedAt:      post.CreatedAt,
            Likes:          likes,
            Dislikes:       dislikes,
        }
	}
	return postsResponse, nil
}

func (h *VoteHandler) AppendVotesToCommentsResponse(comments []model.Comment) ([]model.CommentsResponse, error) {
	commentsResponse := make([]model.CommentsResponse, len(comments))
	for i, comment := range comments {
		likes, dislikes, err := h.voteRepo.GetItemVotes("comment", comment.Id)
		if err != nil {
			return nil, err
		}
		commentsResponse[i] = model.CommentsResponse{
			Id:        comment.Id,
			PostID:    comment.PostID,
			UserID:    comment.UserID,
			Content:   comment.Content,
			Image:     comment.Image.String,
			CreatedAt: comment.CreatedAt,
			Likes:     likes,
			Dislikes:  dislikes,
		}
	}
	return commentsResponse, nil
}