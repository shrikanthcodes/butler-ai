package model

const DEFAULT_LLM = "openai"                // Default language model to use for the conversation
const DEFAULT_LLM_VERSION = "gpt-3.5-turbo" // Default model to use for the conversation

type ChatDialogue struct { //struct to store a single dialogue in a conversation
	Role    string `json:"role"`    // Role of the speaker (system or user or assistant)
	Content string `json:"content"` // Content of the dialogue
}

type ChatConversation struct { //struct to store user's conversation history (Each user can have multiple conversation histories with the assistant)
	ConversationID string         `json:"conversation_id"` // Conversation ID
	UserID         string         `json:"user_id"`         // User ID, foreign key
	Conversation   []ChatDialogue `json:"conversation"`    // A list of dialogues spoken by the user and the assistant
	LastUpdated    string         `json:"last_updated"`    // Last updated timestamp
}
