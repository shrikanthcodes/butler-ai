package templates

// setTemplate returns the templates file string based on the chat type.
func setTemplate(chatType string) string {
	switch chatType {
	case "recipe":
		return "recipe-mode-001.tmpl"
	case "shopping":
		return "shopping-mode-001.tmpl"
	case "health":
		return "health-mode-001.tmpl"
	case "motivation":
		return "motivation-mode-001.tmpl"
	case "calorie_tracker":
		return "calorie-tracker-mode-001.tmpl"
	case "summarization":
		return "chat-summarization-mode-001.tmpl"
	default:
		return ""
	}
}
