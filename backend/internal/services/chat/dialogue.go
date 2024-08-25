package chat

import (
	openai "backend/api/openai"
	model "backend/internal/model"
	bufio "bufio"
	fmt "fmt"
	os "os"
	strings "strings"
)

func GetReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func GetUserMessage(conversation []model.ChatDialogue, reader *bufio.Reader) (bool, []model.ChatDialogue) {
	// Get the user message from the user (in a real application, this would be from the frontend)
	exitCondition := false

	// Get the user message
	fmt.Printf("User: ")
	userMessage, _ := reader.ReadString('\n')
	userMessage = strings.TrimSpace(userMessage)

	// Check if the user wants to exit (temp logic)
	if userMessage == "exit" {
		exitCondition = true
	} else {
		// Update conversation with user message
		conversation = append(conversation, model.ChatDialogue{Role: "user", Content: userMessage})
	}

	return exitCondition, conversation
}

func GetAIMessage(conversation []model.ChatDialogue) []model.ChatDialogue {
	// Find the next message
	ai_model := GetModel()
	response := openai.ChatComplete(ai_model, conversation)
	response_message := string(response.Choices[0].Message.Content)

	// Output the response
	fmt.Printf("Butler: %v\n", response_message)

	// Update conversation with system message
	conversation = append(conversation, model.ChatDialogue{Role: "assistant", Content: response_message})

	return conversation
}
