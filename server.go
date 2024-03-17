package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

var (
	messages  []Message
	mu        sync.Mutex
	broadcast = make(chan struct{}, 100) // Increase buffer size
	pl        = fmt.Println
)

type Message struct {
	DisplayName string `json:"displayName"`
	Message     string `json:"message"`
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Respond with the messages as a JSON array
	json.NewEncoder(w).Encode(messages)
}

func handleNewMessage(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages = append(messages, msg)
	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	pl("New message:", msg.Message, "\nFrom:", msg.DisplayName, "\nIP:", clientIP)

	// Signal that a new message has arrived
	pl("Signaling new message")
	select {
	case broadcast <- struct{}{}:
	default:
		pl("Broadcast channel is full, dropping message")
	}
}

func main() {
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Define other endpoints
	http.HandleFunc("/messages", handleMessages)
	http.HandleFunc("/new-message", handleNewMessage)

	pl("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
