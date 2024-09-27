package models

type Importance struct {
	Importance string   `json:"importance"` // Importance of the item (e.g. "high", "medium", "low")
	Items      []string `json:"items"`      // List of items
}

type BasicInfo struct { //Modular struct to store any data that conforms to basic information structure (name and description). Used for for health conditions, medications, allergies, dietary restrictions, etc
	Name        string `json:"name"`        // Name of the item
	Description string `json:"description"` // Description of the item
}

// FoodItem represents an item in the user's pantry.
type FoodItem struct { //struct to store food item information (Can be used in FoodInventory, or Recipe struct)
	ItemName string  // Name of the item
	Quantity float64 // Quantity of the item
	Unit     string  // Unit of measurement (e.g. "pieces", "grams", "liters")
}

type FoodUnits struct { //struct to store user's preferred units for food items (usually stored at user level)
	Weight       string `json:"weight"`       //string to store weight units (e.g. "kg", "g", "lb", etc)
	Volume       string `json:"volume"`       //string to store volume units (e.g. "litre", "fl. oz", "ml" etc)
	Height       string `json:"height"`       //string to store height units (e.g. "inch", "cm", etc)
	Conventional string `json:"conventional"` //string to store conventional units (e.g. "cup", "tablespoon", etc)
	Colloquial   string `json:"colloquial"`   //string to store colloquial units (e.g. "handful", "pinch", etc)
}

type RecipeInstructions struct { //struct to store recipe instructions (step by step) (usually contains instructions for a recipe, but also for specific tasks like instructions for sauteing onions, blanching tomatoes, etc)
	Step        int    `json:"step"`        // Step number
	Instruction string `json:"instruction"` // Instruction for the step
	Media       string `json:"media"`       // URL to video or image
}

type RecipeTime struct { //struct to store recipe time information (Available at recipe level or user recipe generation level)
	CookingTime     string `json:"cooking_time"`      //string to store cooking time
	CookingTimeUnit string `json:"cooking_time_unit"` //string to store cooking time unit (minutes, hours, etc)
	PrepTime        string `json:"prep_time"`         //string to store prep time
	PrepTimeUnit    string `json:"prep_time_unit"`    //string to store prep time unit (minutes, hours, etc)
	TotalTime       string `json:"total_time"`        //string to store total time
}