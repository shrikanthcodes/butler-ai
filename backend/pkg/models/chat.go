package models

import "log"

type Dialogue struct {
	Role    string `json:"role"`    // Role of the message (e.g. "user", "model")
	Content string `json:"content"` // Content of the message
}

func SetRole(role string) string {
	switch role {
	case "user":
		return "user"
	case "model":
		return "model"
	default:
		log.Fatalf("Invalid role")
		return ""
	}
}
