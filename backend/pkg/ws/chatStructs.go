package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub *Hub
	// Websocket connection
	Conn *websocket.Conn
	// Buffered channel for outbound messages
	Send     chan []byte
	ID       int
	Username string
	Online   bool
}

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from the clients.
	Unregister chan *Client

	ChatHandler *ChatHandler
}

type FetchMessage struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data,omitempty"`
}

type ChatMessage struct {
	MessageID  int    `json:"id"`
	SenderID   int    `json:"sender"`
	ReceiverID int    `json:"receiver"`
	Message    string `json:"text"`
	CreatedAt  string `json:"timestamp"`
}
