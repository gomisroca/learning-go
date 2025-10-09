package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// We will build a simple server that echoes back any messages sent to it.
// We will use the gorilla/websocket package to handle the WebSocket connection.

// Struct used to upgrade HTTP connection to WebSocket connection.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print message to console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	// Serve the HTML page that will listen to /echo.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })
	
    http.ListenAndServe(":8080", nil)
}