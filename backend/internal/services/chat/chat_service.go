package chat

import (
	config "backend/internal/config"
	templates "backend/internal/services/template"
	user "backend/internal/services/user"
	fmt "fmt"
)

// GetModel retrieves the model to use for the conversation
func GetModel() string {
	return config.DEFAULT_MODEL
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
func NewConversation(conversationID string) []config.Dialogue {
	conversation := []config.Dialogue{
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

// SaveConversation saves the conversation state to a database
func SaveConversation(conversationID string, conversation []config.Dialogue) {
	fmt.Printf("Saving conversation with ID %v\n", conversationID)
}
