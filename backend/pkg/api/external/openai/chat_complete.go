package openai

import (
	model "backend/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ChatComplete(ai_model string, api_key string, conversation []model.Dialogue) model.OpenAIResponse {

	url := "https://api.openai.com/v1/chat/completions"

	// Convert the conversation slice to JSON
	conversationJSON, err := json.Marshal(conversation)
	if err != nil {
		fmt.Println("Error marshalling conversation:", err)
		return model.OpenAIResponse{}
	}

	// Construct the payload
	payload := fmt.Sprintf(`{
		"model": "%s",
		"messages": %s
	}`, ai_model, conversationJSON)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return model.OpenAIResponse{}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return model.OpenAIResponse{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return model.OpenAIResponse{}
	}

	var responseBody model.OpenAIResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return model.OpenAIResponse{}
	}

	return responseBody
}
