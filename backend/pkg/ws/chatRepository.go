package ws

import (
	"database/sql"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (h *ChatRepository) GetMessages(senderID, recipientID int, page int) ([]ChatMessage, error) {
	perPage := 10
	offset := (page - 1) * perPage

	// Count total amount of message to know when to stop fetching
	countQuery := "SELECT COUNT(*) FROM chats WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)"
	var totalMessages int
	err := h.db.QueryRow(countQuery, senderID, recipientID, recipientID, senderID).Scan(&totalMessages)
	if err != nil {
		return nil, err
	}

	if offset >= totalMessages {
		return []ChatMessage{}, nil
	} else {
		// Modify your SQL query to limit and offset
		query := "SELECT id, sender_id, receiver_id, message, strftime('%Y-%m-%d %H:%M:%S', created_at, '+3 hours') AS formattedCreatedAt FROM chats WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?) ORDER BY created_at DESC LIMIT ? OFFSET ?"
		rows, err := h.db.Query(query, senderID, recipientID, recipientID, senderID, perPage, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var chatHistory []ChatMessage
		for rows.Next() {
			var msg ChatMessage
			if err := rows.Scan(&msg.MessageID, &msg.SenderID, &msg.ReceiverID, &msg.Message, &msg.CreatedAt); err != nil {
				return nil, err
			}
			chatHistory = append(chatHistory, msg)
		}

		return chatHistory, nil
	}
}

func (h *ChatRepository) StoreMessage(senderID, recipientID int, message string) error {
	_, err := h.db.Exec("INSERT INTO chats (sender_id, receiver_id, message) VALUES (?, ?, ?)", senderID, recipientID, message)
	return err
}
