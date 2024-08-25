package model

type Recipe struct { //struct to store recipe information (Each recipe has a name, category, cuisine, ingredients, instructions, etc)
	RecipeID        string               `json:"recipe_id"`        // Recipe ID
	Time            RecipeTime           `json:"recipe_time"`      // Time required to cook the recipe
	Name            string               `json:"name"`             // Name of the recipe
	Category        string               `json:"category"`         // Category of the recipe
	Cuisine         string               `json:"cuisine"`          // Cuisine of the recipe
	Ingredients     []FoodItem           `json:"ingredients"`      // Ingredients required for the recipe
	Instructions    []RecipeInstructions `json:"instructions"`     // Step by step instructions for the recipe
	NutritionalInfo string               `json:"nutritional_info"` // Nutritional information of the recipe
	UserID          string               `json:"user_id"`          // User ID of the user who created the recipe
	RecipeHTML      string               `json:"recipe_html"`      // HTML content of the recipe
}

type RecipeInstructions struct { //struct to store recipe instructions (step by step) (usually contains instructions for a recipe, but also for specific tasks like instructions for sauteing onions, blanching tomatoes, etc)
	Step        int    `json:"step"`        // Step number
	Instruction string `json:"instruction"` // Instruction for the step
	Media       string `json:"media"`       // URL to video or image
}

type RecipePreferences struct { //struct to store user's recipe preferences (Each user has preferences for recipe generation like serving size, shopping preferences, time available, etc, might change for each recipe so confirm)
	UserID              string              `json:"user_id"`              // User ID
	ServingSize         int                 `json:"serving_size"`         //serving size for the recipe
	ShoppingPreferences ShoppingPreferences `json:"shopping_preferences"` //shopping preferences for the recipe
	TimeAvailable       RecipeTime          `json:"time_available"`       //time available for cooking
	Innovative          bool                `json:"innovative"`           //boolean to store if user is willing to try innovative recipes
}

type RecipeTime struct { //struct to store recipe time information (Available at recipe level or user recipe generation level)
	CookingTime     string `json:"cooking_time"`      //string to store cooking time
	CookingTimeUnit string `json:"cooking_time_unit"` //string to store cooking time unit (minutes, hours, etc)
	PrepTime        string `json:"prep_time"`         //string to store prep time
	PrepTimeUnit    string `json:"prep_time_unit"`    //string to store prep time unit (minutes, hours, etc)
}
