package chat

import (
	openai "backend/api/openai"
	config "backend/config"
	bufio "bufio"
	fmt "fmt"
	os "os"
	strings "strings"
)

func GetReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func GetUserMessage(conversation []config.Dialogue, reader *bufio.Reader) (bool, []config.Dialogue) {
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
		conversation = append(conversation, config.Dialogue{Role: "user", Content: userMessage})
	}

	return exitCondition, conversation
}

func GetAIMessage(conversation []config.Dialogue) []config.Dialogue {
	// Find the next message
	model := GetModel()
	response := openai.ChatComplete(model, conversation)
	response_message := string(response.Choices[0].Message.Content)

	// Output the response
	fmt.Printf("Butler: %v\n", response_message)

	// Update conversation with system message
	conversation = append(conversation, config.Dialogue{Role: "assistant", Content: response_message})

	return conversation
}
