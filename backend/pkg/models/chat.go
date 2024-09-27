package models

// These are supplementary data structures (not a table in DB)
type Dialogue struct {
	Role    string `json:"role"`    // Role of the message (e.g. "system", "user", "assistant")
	Content string `json:"content"` // Content of the message
}

type Role string // Role of the message (e.g. "system", "user", "assistant")

const (
	SystemRole    Role = "system"    // System role for the AI
	UserRole      Role = "user"      // User role for the AI
	AssistantRole Role = "assistant" // Assistant role for the AI
)
