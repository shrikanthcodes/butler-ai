package chat

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
