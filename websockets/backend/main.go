package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var clients = make(map[*websocket.Conn]bool) // ⬅️ Map of connected clients
var mu sync.Mutex // Prevent concurrent access to clients map

var upgrader = websocket.Upgrader{ // ⬅️ Upgrader for WebSocket (think of as connection)
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow any origin (not secure for production)
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // ⬅️ Upgrade HTTP to WebSocket
    if err != nil {
        fmt.Println("Upgrade error:", err)
        return
    }
    defer conn.Close() // ⬅️ Make sure to close connection when done

	// Add client
	mu.Lock() // ⬅️ Protect access to clients map
	clients[conn] = true // ⬅️ Add client to map
	mu.Unlock() // ⬅️ Unlock access to clients map

	fmt.Println("New client connected")

	// Remove client on disconnect
	defer func() {
		mu.Lock() // ⬅️ Protect access to clients map
		delete(clients, conn) // ⬅️ Remove client from map
		mu.Unlock() // ⬅️ Unlock access to clients map
		fmt.Println("Client disconnected")
	}()


   	for {
		_, msgBytes, err := conn.ReadMessage() // ⬅️ Read message from client
		if err != nil {
			break
		}

		var msg Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil { // ⬅️ Convert JSON to Message struct
			fmt.Println("Error unmarshalling message:", err) 
			continue
		}

		// Broadcast to all clients
		broadcast(msg)
	}
}

func broadcast(msg Message) {
	msgBytes, err := json.Marshal(msg) // ⬅️ Convert Message struct to JSON
	if err != nil {
		return
	}


	mu.Lock() // ⬅️ Protect access to clients map
	defer mu.Unlock() // ⬅️ Unlock access to clients map when done
	
	for client := range clients { // ⬅️ Iterate over clients map
		if err := client.WriteMessage(websocket.TextMessage, msgBytes); err != nil { // ⬅️ Write message to client
			client.Close()
			delete(clients, client) // ⬅️ Remove client from map when error
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
