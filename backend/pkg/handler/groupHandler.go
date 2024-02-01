package handler

import (
	"backend/pkg/db/sqlite"
	"backend/pkg/model"
	"backend/pkg/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Group Handlers
func GetAllGroupsHandler(w http.ResponseWriter, r *http.Request) {
    // check auth and get userid from cookie
    cookie, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // If the session cookie doesn't exist, set isAuthenticated to false
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        } else {
            http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
            return
        }
    }
    _, err = confirmAuthentication(cookie) //(userid, err :=) when later would highlight groups user is a part of
    if err != nil {
        http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusUnauthorized)
        return
    }
    // TODO: Implement logic for getting all groups
    groups, err := repository.NewGroupRepository(sqlite.Dbase).GetAllGroups()
    if err != nil {
        http.Error(w, "Failed to get groups: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(groups)
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for creating a group
    var newGroup model.Group
    err := json.NewDecoder(r.Body).Decode(&newGroup)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    // check auth and get userid from cookie
    cookie, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // If the session cookie doesn't exist, set isAuthenticated to false
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        } else {
            http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
            return
        }
    }
    userid, err := confirmAuthentication(cookie)
    if err != nil {
        http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusUnauthorized)
        return
    }
    // TODO: check if group with title already exists

    newGroup.CreatorId = userid
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    createdGroup, err := groupRepo.CreateGroup(newGroup)
    if err != nil {
        http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(createdGroup)

}

func GetGroupByIDHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for getting a group by ID
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    group, err := groupRepo.GetGroupByID(id)
    if err != nil {
        http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(group)
}

func EditGroupHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement logic for editing a group
    var updatedGroup model.Group
    err := json.NewDecoder(r.Body).Decode(&updatedGroup)
    if err != nil {
        http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    err = groupRepo.UpdateGroup(updatedGroup)
    if err != nil {
        http.Error(w, "Failed to update group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedGroup)
}

// DeleteGroupHandler handles the HTTP request for deleting a group.
// It checks the user's authentication, verifies their authorization to delete the group,
// and deletes the group from the repository if all conditions are met.
// If any errors occur during the process, appropriate HTTP error responses are returned.
// TODO: implement notification to all group members that the group has been deleted, and remove all group members;
// implement logging of the deletion or add bool field "deleted"
func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
    // check auth and get userid from cookie
    cookie, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // If the session cookie doesn't exist, set isAuthenticated to false
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        } else {
            http.Error(w, "Error checking session token: " + err.Error(), http.StatusInternalServerError)
            return
        }
    }
    userId, err := confirmAuthentication(cookie)
    if err != nil {
        http.Error(w, "Error confirming authentication: " + err.Error(), http.StatusUnauthorized)
        return
    }

    // TODO: Implement logic for deleting a group
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }
    groupRepo := repository.NewGroupRepository(sqlite.Dbase)
    group, err := groupRepo.GetGroupByID(id)
    if err != nil {
        http.Error(w, "Failed to get group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    if group.CreatorId != userId {
        http.Error(w, "User not authorized to delete this group", http.StatusUnauthorized)
        return
    }
    err = groupRepo.DeleteGroup(id)
    if err != nil {
        http.Error(w, "Failed to delete group: "+err.Error(), http.StatusInternalServerError)
        return
    }
    // Successful response
	response := map[string]string{
		"message": "Post deleted successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}