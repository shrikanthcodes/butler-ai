package chat

import (
	config "backend/config"
	templates "backend/services/template"
	fmt "fmt"
)

// GetModel retrieves the model to use for the conversation
func GetModel() string {
	return config.DEFAULT_MODEL
}

// GetUserData retrieves user data based on the user ID
func GetUserData(userID string) config.RecipeData {
	pantryItems := []config.PantryItem{
		{ItemName: "Tomatoes", Quantity: 5, Unit: "pieces"},
		{ItemName: "Onions", Quantity: 3, Unit: "pieces"},
		{ItemName: "Garlic", Quantity: 2, Unit: "cloves"},
	}

	data := config.RecipeData{
		Name:        "Aishwarya",
		Allergy:     "Peanuts",
		Preferences: "Vegetarian",
		PantryItems: pantryItems,
	}
	return data
}

// GetUserPrompt generates a user prompt based on the user's data
func GetUserPrompt(userID string) string {
	result, err := templates.RenderRecipeTemplate(GetUserData(userID))
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
