package model

type Dialogue struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

const DEFAULT_MODEL = "gpt-3.5-turbo"

type Response struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	Index        int         `json:"index"`
	Message      Dialogue    `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type RecipeData struct {
	Name        string
	Allergy     string
	Preferences string
	PantryItems []PantryItem
}

// PantryItem represents an item in the user's pantry.
type PantryItem struct {
	ItemName string
	Quantity float64
	Unit     string
}
