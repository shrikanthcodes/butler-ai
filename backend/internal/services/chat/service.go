package chat

import (
	model "backend/internal/model"
	templates "backend/internal/services/template"
	user "backend/internal/services/user"
	bufio "bufio"
	json "encoding/json"
	fmt "fmt"
	http "net/http"
	strings "strings"
)

// GetModel retrieves the model to use for the conversation
func GetModel() string {
	return model.DEFAULT_LLM_VERSION
}

// GetUserPrompt generates a user prompt based on the user's data
func GetUserPrompt(userID string) string {
	result, err := templates.RenderRecipeTemplate(user.GetCompleteUserData(userID))
	if err != nil {
		fmt.Printf("Error rendering template: %v\n", err)
		return ""
	}
	return result
}

// NewConversation initializes a new conversation
func NewConversation(conversationID string) []model.ChatDialogue {
	conversation := []model.ChatDialogue{
		{Role: "system", Content: GetUserPrompt("test")},
	}
	return conversation
}

// Logic to handle back and forth conversations and to save the conversation state
func HandleConversation(conversationID string) {
	conversation, reader := NewConversation(conversationID), GetReader()
	endConversation := false
	for i := 0; i < 5; i++ {
		endConversation, conversation = GetUserMessage(conversation, reader)

		if endConversation {
			break
		}

		conversation = GetAIMessage(conversation)
	}
	SaveConversation(conversationID, conversation)
}

// HandleChatConversation is the HTTP handler for processing conversations
func HandleChatConversation(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get user message and conversation ID
	var req struct {
		ConversationID string               `json:"conversation_id"`
		Messages       []model.ChatDialogue `json:"messages"`
	}

	// Decode the JSON request payload
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize the conversation with the provided messages
	conversation := req.Messages

	// Process the conversation (simulate the back-and-forth)
	for i := 0; i < 5; i++ { // Example loop count for conversation turns
		// Get user message from the request
		endConversation, updatedConversation := GetUserMessage(conversation, bufio.NewReader(strings.NewReader(req.Messages[len(req.Messages)-1].Content)))
		conversation = updatedConversation

		if endConversation {
			break
		}

		// Get AI response and update conversation
		conversation = GetAIMessage(conversation)
	}

	// Save the updated conversation
	SaveConversation(req.ConversationID, conversation)

	// Respond with the updated conversation
	response := struct {
		Messages []model.ChatDialogue `json:"messages"`
	}{
		Messages: conversation,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// SaveConversation saves the conversation state to a database
func SaveConversation(conversationID string, conversation []model.ChatDialogue) {
	fmt.Printf("Saving conversation with ID %v\n", conversationID)
}
