package repository

import (
	"backend/pkg/model"
	"database/sql"
	"fmt"
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

func (r *VoteRepository) GetItemVotes(itemType string, itemID int) (int, int, error){
	var likes, dislikes int
	query := fmt.Sprintf("SELECT COUNT(*) FROM votes WHERE %sID = ? AND type = ?", itemType)
	err := r.db.QueryRow(query, itemID, "upvote").Scan(&likes)
	if err != nil {
		return 0, 0, err
	}
	err = r.db.QueryRow(query, itemID, "downvote").Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}