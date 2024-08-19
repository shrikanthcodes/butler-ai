package chat

import (
	config "backend/config"
	templates "backend/services/template"
	bufio "bufio"
	fmt "fmt"
	os "os"
	strings "strings"
)

func GetModel() string {
	return config.DEFAULT_MODEL
}

func GetUserData(userID string) config.RecipeData {
	// Example pantry items
	pantryItems := []config.PantryItem{
		{ItemName: "Tomatoes", Quantity: 5, Unit: "pieces"},
		{ItemName: "Onions", Quantity: 3, Unit: "pieces"},
		{ItemName: "Garlic", Quantity: 2, Unit: "cloves"},
	}

	// Create a data instance to pass to the template
	data := config.RecipeData{
		Name:        "Aishwarya",
		Allergy:     "Peanuts",
		Preferences: "Vegetarian",
		PantryItems: pantryItems,
	}
	return data
}

func GenerateUserPrompt(userID string) string {

	// Use the dedicated function to render the recipe template
	result, err := templates.RenderRecipeTemplate(GetUserData(userID))
	if err != nil {
		fmt.Printf("Error rendering template: %v\n", err)
		return ""
	}

	// Output the result
	return result
}

// GetConversation returns a slice of Dialogue structs, takes in a conversationID string as input, if not provided, defaults to "test"
func GetConversation(conversationID string) []config.Dialogue {
	conversation := []config.Dialogue{}
	// Add conversation logic here
	if conversation != nil {
		conversation = []config.Dialogue{
			{Role: "system", Content: GenerateUserPrompt("test")},
		}
	}

	return conversation
}

// Logic to handle back and forth conversations and to save the conversation state
func HandleConversation(conversationID string) {
	// Get the model and conversation until now
	conversation, model := GetConversation(conversationID), GetModel()

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 5; i++ {
		// Get the user message from the user (in a real application, this would be from the frontend)
		fmt.Printf("User: ")
		userMessage, _ := reader.ReadString('\n')
		userMessage = strings.TrimSpace(userMessage)

		if userMessage == "exit" {
			break
		}

		// Update conversation with user message
		conversation = append(conversation, config.Dialogue{Role: "user", Content: userMessage})

		// Find the next message
		response := NextMessage(model, conversation)
		fmt.Printf("Butler: %v\n", response)

		// Update conversation with system message
		conversation = append(conversation, config.Dialogue{Role: "assistant", Content: response})
	}

	// Save the conversation at the end
	SaveConversation(conversationID, conversation)
}

// SaveConversation saves the conversation state to a database or file
func SaveConversation(conversationID string, conversation []config.Dialogue) {
	fmt.Printf("Saving conversation with ID %v\n", conversationID)
	// Save the conversation state
}
