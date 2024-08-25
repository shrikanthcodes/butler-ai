package model

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

type FoodInventory struct { //struct to store user's inventory (Each user has an inventory of food items)
	UserID string     `json:"inventory_id"` // User ID
	Items  []FoodItem `json:"items"`        // FoodItem struct is defined in openai.go
}
