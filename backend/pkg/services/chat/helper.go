package chat

import (
	models "backend/pkg/models"
	templates "backend/pkg/services/templates"
	errors "errors"
)

// getTemplateNameByChatType returns the template name based on the chat type.
func getTemplateNameByChatType(chatType string) (string, error) {
	switch chatType {
	case "recipe":
		return templates.RECIPE_MODE, nil
	case "shopping":
		return templates.SHOPPING_MODE, nil
	case "health":
		return templates.HEALTH_MODE, nil
	case "motivation":
		return templates.MOTIVATION_MODE, nil
	case "calorie_tracker":
		return templates.CALORIE_TRACKER_MODE, nil
	case "summarization":
		return templates.SUMMARIZATION_MODE, nil
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
}

const MaxConversationLength = 10000 // Maximum length before summarization.

// calculateContextLength calculates the total length of the conversation.
func calculateContextLength(dialogues []models.Dialogue) int {
	totalLength := 0
	for _, d := range dialogues {
		totalLength += len(d.Content)
	}
	return totalLength
}
