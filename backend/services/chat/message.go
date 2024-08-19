package chat

import (
	openai "backend/api/openai"
	config "backend/config"
)

func NextMessage(model string, conversation []config.Dialogue) string {
	response := openai.ChatComplete(model, conversation)
	//println(response.Choices[0].Message.Content)
	return string(response.Choices[0].Message.Content)
}
