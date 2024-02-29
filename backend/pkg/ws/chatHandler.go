package ws

import (
	"backend/pkg/repository"
	"fmt"
	"log"
	"time"
)

type ChatHandler struct {
	ChatRepo    *ChatRepository
	SessionRepo *repository.SessionRepository
}

func NewChatHandler(chatRepo *ChatRepository, sessionRepo *repository.SessionRepository) *ChatHandler {
	return &ChatHandler{ChatRepo: chatRepo, SessionRepo: sessionRepo}
}
func (h *ChatHandler) FetchChatHistory(c *Client, recipientID int, page int) {

	chatHistory, err := h.ChatRepo.GetMessages(c.ID, recipientID, page)
	if err != nil {
		log.Printf("Error fetching chat history: %v", err)
		return
	}

	response := make(map[string]interface{})
	response["action"] = "chat_history"
	response["content"] = chatHistory

	c.Conn.WriteJSON(response)
}

func (h *ChatHandler) SendMessage(messageData map[string]interface{}, c *Client) {
	message, ok := messageData["content"].(string)
	if !ok {
		log.Printf("Invalid message format: %v", messageData)
	}
	if c == nil {
		fmt.Println("client nil")
	}
	messageData["timestamp"] = time.Now().Format(time.RFC3339)
	messageData["sender"] = c.ID
	recipientID := int(messageData["recipientID"].(float64))
	for key, value := range c.Hub.Clients {
		if key.ID == recipientID {
			if value {
				key.Conn.WriteJSON(messageData)
			}
		}
	}

	err := h.ChatRepo.StoreMessage(c.ID, recipientID, message)
	if err != nil {
		log.Print("Error while storing message to database")
		return
	}

}
