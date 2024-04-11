package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
	"log" // Import the log package for logging
)

type VoteRepository struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) *VoteRepository {
	return &VoteRepository{db: db}
}

func (r *VoteRepository) VoteItem(vote model.VoteData, userID int) error {
	// Check if the user has already rated the item
	var currentVote string
	query := fmt.Sprintf("SELECT type FROM votes WHERE %sID = ? AND userID = ?", vote.Item)
	err := r.db.QueryRow(query, vote.ItemID, userID).Scan(&currentVote)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = r.db.Exec(fmt.Sprintf("INSERT INTO votes (%sID, userID, type) VALUES (?, ?, ?)", vote.Item), vote.ItemID, userID, vote.Action)
			fmt.Println(err)
			return err
		}
		return err
	}
	if vote.Action == currentVote {
		// If the user has already rated the item with the same rating, remove the rating
		_, err = r.db.Exec(fmt.Sprintf("DELETE FROM votes WHERE %sID = ? AND userID = ?", vote.Item), vote.ItemID, userID)
		return err
	}

	// If the user has rated the item differently, update the rating
	_, err = r.db.Exec(fmt.Sprintf("UPDATE votes SET type = ? WHERE %sID = ? AND userID = ?", vote.Item), vote.Action, vote.ItemID, userID)
	return err
}

func (r *VoteRepository) GetItemVotes(itemType string, itemID int) (int, int, error) {
	var likes, dislikes int
	query := fmt.Sprintf("SELECT COUNT(*) FROM votes WHERE %sID = ? AND type = ?", itemType)
	err := r.db.QueryRow(query, itemID, "like").Scan(&likes)
	if err != nil {
		log.Printf("Error retrieving upvotes for %s ID %d: %v\n", itemType, itemID, err)
		return 0, 0, err
	}
	err = r.db.QueryRow(query, itemID, "dislike").Scan(&dislikes)
	if err != nil {
		log.Printf("Error retrieving downvotes for %s ID %d: %v\n", itemType, itemID, err)
		return 0, 0, err
	}
	return likes, dislikes, nil
}

func (r *VoteRepository) GetUserVoteAction(userID int, itemType string, itemID int) (string, error) {
	var action string
	query := fmt.Sprintf("SELECT type FROM votes WHERE %sID = ? AND userID = ?", itemType)
	err := r.db.QueryRow(query, itemID, userID).Scan(&action)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no vote action found for the user and item, return an empty string
			return "", nil
		}
		log.Printf("Error retrieving user vote action for %s ID %d: %v\n", itemType, itemID, err)
		return "", err
	}
	return action, nil
}

/* func (r *VoteRepository) HasUserVoted(itemType string, itemID, userID int) bool {
	// Check if the user has voted for the given item
	var voteCount int
	query := fmt.Sprintf("SELECT COUNT(*) FROM votes WHERE %sID = ? AND userID = ?", itemType)
	err := r.db.QueryRow(query, itemID, userID).Scan(&voteCount)
	if err != nil {
		log.Printf("Error checking user vote for %s ID %d: %v\n", itemType, itemID, err)
		return false // Assume no vote has been cast in case of an error
	}
	return voteCount > 0 // Return true if the user has voted (voteCount > 0), false otherwise
} */
