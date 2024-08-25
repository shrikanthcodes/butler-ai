package main

import (
	chat "backend/internal/services/chat"
	http "net/http"

	mux "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define the route for handling chat conversations
	router.HandleFunc("/chat/conversation", chat.HandleChatConversation).Methods("POST")

	// Start the server
	http.ListenAndServe(":8080", router)
}
