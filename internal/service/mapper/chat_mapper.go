package mapper

import "errors"

// setTemplate returns the template file string based on the chat type.
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

func MapAIParametersToChatType(chatType string) (string, int32, float32, error) {
	switch chatType {
	case "recipe":
		return setTemplate("recipe"), setResponseLength("long"), setTemperature("creative"), nil
	case "shopping":
		return setTemplate("shopping"), setResponseLength("short"), setTemperature("regular"), nil
	case "health":
		return setTemplate("health"), setResponseLength("long"), setTemperature("regular"), nil
	case "motivation":
		return setTemplate("motivation"), setResponseLength("medium"), setTemperature("creative"), nil
	case "calorie_tracker":
		return setTemplate("calorie_tracker"), setResponseLength("medium"), setTemperature("deterministic"), nil
	case "summarization":
		return setTemplate("summarization"), setResponseLength("long"), setTemperature("regular"), nil
	default:
		return "", setResponseLength("medium"), setTemperature("regular"), errors.New("invalid chat type")
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
}
