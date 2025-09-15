package controller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins for now
	},
}

// store connected clients
var clients = make(map[string]*websocket.Conn)
var mu sync.Mutex

func ChatHandler(c *gin.Context) {
	userID := c.Query("user_id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	// add client
	mu.Lock()
	clients[userID] = conn
	mu.Unlock()
	fmt.Println("✅ User", userID, "connected")

	defer func() {
		// remove client on disconnect
		mu.Lock()
		delete(clients, userID)
		mu.Unlock()
		conn.Close()
		fmt.Println("❌ User", userID, "disconnected")
	}()

	// listen for messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("User %s says: %s\n", userID, string(message))

		// broadcast message to all connected clients
		mu.Lock()
		for id, client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("User %s: %s", userID, message))); err != nil {
				fmt.Println("Error writing to client", id, ":", err)
			}
		}
		mu.Unlock()
	}
}
