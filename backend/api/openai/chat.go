package openai

import (
	model "backend/internal/model"
	secrets "backend/secrets"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ChatComplete(ai_model string, conversation []model.Dialogue) model.Response {
	// Make a POST request to the OpenAI API
	api_key := secrets.OpenAI_Key()

	url := "https://api.openai.com/v1/chat/completions"

	// Convert the conversation slice to JSON
	conversationJSON, err := json.Marshal(conversation)
	if err != nil {
		fmt.Println("Error marshalling conversation:", err)
		return model.Response{}
	}

	// Construct the payload
	payload := fmt.Sprintf(`{
		"model": "%s",
		"messages": %s
	}`, ai_model, conversationJSON)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return model.Response{}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return model.Response{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return model.Response{}
	}

	var responseBody model.Response
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return model.Response{}
	}

	return responseBody
}
