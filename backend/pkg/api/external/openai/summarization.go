package openai

import (
	model "backend/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: Add logic to catch if user's usage limit is exceeded, if yes, try to find and use another model in their repo if exists
// TODO: Trigger automnatic summarization of past conversations (add to initial prompt) once conversation exceeds a certain length keeping only the 3 most recent to-and-fros
func Summarization(ai_model string, api_key string, text model.Dialogue, max_tokens int, temperature float32) (model.OpenAIResponse, error) {

	url := "https://api.openai.com/v1/chat/completions"

	// Convert the conversation slice to JSON
	textJSON, err := json.Marshal(text)
	if err != nil {
		log.Fatalf("Error marshalling conversation: %s", err)
	}

	// Check if any of the input is invalid
	if text.Content == "" || ai_model == "" || api_key == "" || max_tokens == 0 {
		log.Fatalf("Wrong function parameters for summarization, check again")
	}

	// Construct the payload
	payload := fmt.Sprintf(`{
		"model": "%s",
		"messages": "[{"role":"user", "content":%s}]",
		"temperature" = %f,
		"max_tokens" = %d
	}`, ai_model, textJSON, temperature, max_tokens)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Returns to catch Error 429 so it can be handled gracefully
		log.Printf("Error making request: %s", err)
		return model.OpenAIResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}

	var responseBody model.OpenAIResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %s", err)
	}

	return responseBody, err
}
