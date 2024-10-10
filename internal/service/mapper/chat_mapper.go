package mapper

import "errors"

// MapTemplateToChatType returns the template file string based on the chat type.
func MapTemplateToChatType(chatType string) (string, error) {
	switch chatType {
	case "recipe":
		return "recipe-mode-001.tmpl", nil
	case "shopping":
		return "shopping-mode-001.tmpl", nil
	case "health": return "health-mode-001.tmpl", nil
	case "motivation":
		return "motivation-mode-001.tmpl", nil
	case "calorie_tracker":
		return "calorie-tracker-mode-001.tmpl", nil
	case "summarization": return "chat-summarization-mode-001.tmpl", nil
	default:
		return "", errors.New("invalid chat type")
	}
}

func setParametersByChatType(chatType string) (int32, float32) {
	switch chatType {
	case "recipe":
		return setResponseLength("long"), setTemperature("creative")
	case "shopping":
		return setResponseLength("short"), setTemperature("regular")
	case "health":
		return setResponseLength("long"), setTemperature("regular")
	case "motivation":
		return setResponseLength("medium"), setTemperature("creative")
	case "calorie_tracker":
		return setResponseLength("medium"), setTemperature("deterministic")
	case "summarization":
		return setResponseLength("long"), setTemperature("regular")
	default:
		return setResponseLength("medium"), setTemperature("regular")
	}
}

func setTemperature(temperature string) float32 {
	switch temperature {
	case "creative":
		return 0.8
	case "deterministic":
		return 0.2
	case "regular":
		return 0.5
	default:
		return 0.5
	}
}

func setResponseLength(outputLength string) int32 {
	switch outputLength {
	case "short":
		return 250
	case "medium":
		return 500
	case "long":
		return 1200
	default:
		return 1000
	}