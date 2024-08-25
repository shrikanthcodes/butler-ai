package model

type DietaryPreferences struct { //struct to store user's preferences (User level/Each user has preferences for recipes, food items, cuisines, etc)
	UserID          string    `json:"user_id"`             // User ID
	Preferred_Units FoodUnits `json:"preferred_units"`     // Food units preferred by user (units for weight, volume, height, etc)
	FavoriteRecipes []string  `json:"favorite_recipes"`    // Recipes liked by user
	DislikedRecipes []string  `json:"disliked_recipes"`    // Recipes disliked by user
	FavoriteItems   []string  `json:"favorite_items"`      // Items liked by user
	DislikedItems   []string  `json:"disliked_items"`      // Items disliked by user
	FavoriteCuisine []string  `json:"favorite_categories"` // Cuisines of food liked by user (cuisines: Italian, Mexican, etc)
	DislikedCuisine []string  `json:"disliked_categories"` // Cuisines of food disliked by user (cuisines: Italian, Mexican, etc)
}

type ShoppingPreferences struct { //struct to store user's shopping preferences (usually stored at user level and confirmed for each recipe)
	WillingnessToShop     bool   `json:"willingness_to_shop"`     //boolean to store willingness to shop
	BudgetCurrency        string `json:"budget_currency"`         //string to store currency (USD, EUR, etc)
	Budget                string `json:"budget"`                  //string to store budget amount
	EaseOfAvailability    string `json:"ease_of_availability"`    //options are enum: "hard", "medium", "easy"
	ShoppingTimeAvailable string `json:"shopping_time_available"` //string to store time available for shopping
}
