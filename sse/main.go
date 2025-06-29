package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 👇 Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 🥸 Flush allows streaming data in real-time
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// 👇 Stream data every second
	for { // 👈 Infinite loop
		// The EventSource JS API will take this "data: content\n\n" and parse it
		// So the event.data will be f.e "2025-06-29T15:45:00Z"
		fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.RFC3339))
		flusher.Flush() // 👈 Flush data to client
		time.Sleep(1 * time.Second)

		// Check if client has disconnected
		select {
		case <-r.Context().Done():
			log.Println("Client disconnected")
			return
		default:
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	// Serve static HTML
	http.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
