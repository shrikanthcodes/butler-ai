package models

// Preferences for ThingsToConsider (modular JSONB field for flexible user data)
type Preferences struct {
	HealthConditions    []BasicInfo `json:"health_conditions,omitempty"`    // List of health conditions (e.g., diabetes, hypertension)
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions,omitempty"` // List of dietary restrictions (e.g., vegan, gluten-free)
	Allergies           []BasicInfo `json:"allergies,omitempty"`            // List of allergies (e.g., peanuts, shellfish)
	Medications         []BasicInfo `json:"medications,omitempty"`
	FitnessGoals        []string    `json:"fitness_goals,omitempty"` // List of fitness goals (e.g., "Lose 5kg", "Build muscle")

	// Recipe Preferences
	FavoriteRecipes []string `json:"favorite_recipes,omitempty"` // User's favorite recipes by ID or name
	DislikedRecipes []string `json:"disliked_recipes,omitempty"` // Recipes the user dislikes
	FavoriteCuisine []string `json:"favorite_cuisine,omitempty"` // List of user's favorite cuisines (e.g., Italian, Mexican)
	DislikedCuisine []string `json:"disliked_cuisine,omitempty"` // List of cuisines the user dislikes
	FavoriteItems   []string `json:"favorite_items,omitempty"`
	DislikedItems   []string `json:"disliked_items,omitempty"`
	MealPreferences []string `json:"meal_preferences,omitempty"` // List of meal preferences (e.g., high protein, low carb, keto)
	ServingSize     int      `json:"serving_size,omitempty"`     // Preferred serving size for recipes

	// Inventory Preferences
	Inventory []FoodItem `json:"inventory,omitempty"` // Inventory items stored as JSONB

	// Shopping Preferences
	BudgetCurrency    string  `json:"budget_currency,omitempty"`     // Preferred currency for budgeting (e.g., USD, EUR)
	BudgetAmount      float64 `json:"budget_amount,omitempty"`       // Monthly or weekly budget for shopping
	WillingnessToShop bool    `json:"willingness_to_shop,omitempty"` // Willingness to shop for new ingredients
	ShoppingEase      string  `json:"shopping_ease,omitempty"`       // How easy it is for the user to access stores (e.g., "easy", "medium", "hard")
	ShoppingTime      string  `json:"shopping_time,omitempty"`       // Time available for shopping per week (e.g., "2 hours")

	// Nutritional Preferences
	CaloricIntakeGoal float64            `json:"caloric_intake_goal,omitempty"` // Target daily caloric intake
	MacroGoals        map[string]float64 `json:"macro_goals,omitempty"`         // Target macros (e.g., {"protein": 150, "carbs": 200, "fat": 60})

	// User Activity Preferences
	LifestyleActivity string   `json:"lifestyle_activity,omitempty"` // User's lifestyle (e.g., sedentary, active, very active)
	ExerciseRoutine   []string `json:"exercise_routine,omitempty"`   // List of exercises or activities user engages in

	// Other preferences
	CookingExperience string    `json:"cooking_experience,omitempty"` // Level of cooking experience (e.g., beginner, intermediate, advanced)
	PreferredUnits    FoodUnits `json:"preferred_units,omitempty"`    // Preferred units for measurement (e.g., metric, imperial)
	FoodIntolerances  []string  `json:"food_intolerances,omitempty"`  // List of food intolerances (e.g., lactose, gluten)

	// Custom fields for additional user-defined preferences
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"` // Flexible field for arbitrary data
}
