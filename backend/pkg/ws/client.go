package ws

import (
	"backend/util"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var newline = []byte{'\n'}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// TODO: Remove after development
		// Check if the origin is allowed
		return r.Header.Get("Origin") == os.Getenv("NEXT_PUBLIC_URL") + ":" + os.Getenv("NEXT_PUBLIC_HTTPS_PORT")
	},
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.Hub.DisconnectedUserWsAlert(c.ID)
		c.Hub.Unregister <- c
		c.Conn.Close()
		log.Println(c.Username, " disconnected")

	}()
	for {
		_, message, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error : %v", err)
			}
			break
		}
		// Assuming your messages are in JSON format
		var messageData map[string]interface{}
		if err := json.Unmarshal(message, &messageData); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			continue
		}

		// Check the action field in the message
		action, ok := messageData["action"].(string)
		if !ok {
			log.Printf("Invalid message format: %v", messageData)
			continue
		}

		// Handle different actions here
		switch action {

		case "send_message":
			c.Hub.ChatHandler.SendMessage(messageData, c)
		case "fetch_chat_history":
			// Convert the "user" value to an integer
			userID, userOK := messageData["user"].(float64)

			pageInt, pageOK := messageData["page"].(float64)
			if !userOK || !pageOK {
				// Handle the error if any of these conversions fail
				log.Printf("Invalid or missing parameters: userOK=%v, pageOK=%v, perPageOK=%v", userOK, pageOK)
				continue
			}
			c.Hub.ChatHandler.FetchChatHistory(c, int(userID), int(pageInt))
		default:
			// Handle other actions as needed
		}

	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func (h *Hub) ServeWs(w http.ResponseWriter, r *http.Request) {
	userID, err := h.ChatHandler.SessionRepo.GetUserIDFromSessionToken(util.GetSessionToken(r))
	if err != nil {
		log.Println("Error confirming authentication: ", err)
		return
	}
	log.Println("UserID ", userID, " connected")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Hub: h, Conn: conn, Send: make(chan []byte, 256), ID: userID, Online: true}
	h.Register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func (h *Hub) NewUserWsAlert(newUserID int) {
	message := map[string]interface{}{
		"action": "newUser",
		"data":   newUserID,
	}
	jsonData, _ := json.Marshal(message)
	// Broadcast the message to all connected clients
	h.Broadcast <- jsonData
}

func (h *Hub) DisconnectedUserWsAlert(newUserID int) {
	message := map[string]interface{}{
		"action": "disconnectUser",
		"data":   newUserID,
	}
	jsonData, _ := json.Marshal(message)
	// Broadcast the message to all connected clients
	h.Broadcast <- jsonData
}
