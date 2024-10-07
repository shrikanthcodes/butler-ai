package models

type Dialogue struct {
	Role    string `json:"role"`    // Role of the message (e.g. "user", "model")
	Content string `json:"content"` // Content of the message
}

const (
	RoleUser  = "user"
	RoleModel = "model"
)
