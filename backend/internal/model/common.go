package model

type BasicInfo struct { //Modular struct to store any data that conforms to basic information structure (name and description). Used for for health conditions, medications, allergies, dietary restrictions, etc
	Name        string `json:"name"`        // Name of the item
	Description string `json:"description"` // Description of the item
}
