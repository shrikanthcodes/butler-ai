package models

import (
	"gorm.io/gorm"
)

// 1. User
type User struct {
	gorm.Model
	UserID         string `json:"user_id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Email          string `json:"email" gorm:"unique"`
	ProfileWriteup string `json:"profile_writeup"` // LLM generated Profile writeup for the user based on ThingsToConsider (generated each time there is a change in preferences max 1000 len)

	// Relationships to other tables
	Auth         *Authentication `gorm:"foreignKey:UserID"` // Optional relationship
	Choices      *Choice         `gorm:"foreignKey:UserID"` // Optional relationship
	Conversation []Conversation  `gorm:"foreignKey:UserID"` // Optional relationship
	Profile      *Profile        `gorm:"foreignKey:UserID"` // Optional relationship
	Health       *Health         `gorm:"foreignKey:UserID"` // Optional relationship
	Diet         *Diet           `gorm:"foreignKey:UserID"` // Optional relationship
	Inventory    *Inventory      `gorm:"foreignKey:UserID"` // Optional relationship
	Goal         *Goal           `gorm:"foreignKey:UserID"` // Optional relationship
	Script       *Script         `gorm:"foreignKey:UserID"` // Optional relationship
	Shopping     *Shopping       `gorm:"foreignKey:UserID"` // Optional relationship
	MealChoices  *MealChoice     `gorm:"foreignKey:UserID"` // Optional relationship
}

type Choice struct {
	gorm.Model
	UserID string `json:"user_id" gorm:"primaryKey"` // Foreign key
	// Optional Booleans to track which profiles are required
	IsProfile     bool `json:"profile" gorm:"default:false"`
	IsHealth      bool `json:"health" gorm:"default:false"`
	IsDiet        bool `json:"diet" gorm:"default:false"`
	IsInventory   bool `json:"inventory" gorm:"default:false"`
	IsGoal        bool `json:"goal" gorm:"default:false"`
	IsScript      bool `json:"script" gorm:"default:false"`
	IsShopping    bool `json:"shopping" gorm:"default:false"`
	IsPreferences bool `json:"preferences" gorm:"default:false"`
}

// 2. Conversation
type Conversation struct {
	gorm.Model
	ConvID       string     `json:"conv_id" gorm:"primaryKey"`                       // Primary key
	UserID       string     `json:"user_id"`                                         // Foreign key
	User         User       `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Define foreign key
	Conversation []Dialogue `json:"conversation" gorm:"type:jsonb"`                  // Actual conversation between user and AI
	LastUpdated  string     `json:"last_updated"`                                    // Last updated timestamp
	Summary      string     `json:"summary"`
	Task         string     `json:"task"`
	IsActive     bool       `json:"is_active"`
}

// 3. Health
type Health struct {
	gorm.Model
	UserID              string      `json:"user_id"`                                         // Foreign key
	User                User        `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	HealthConditions    []BasicInfo `json:"health_conditions" gorm:"type:jsonb"`             // Health conditions stored as JSONB
	Medications         []BasicInfo `json:"medications" gorm:"type:jsonb"`                   // Medications stored as JSONB
	Allergies           []BasicInfo `json:"allergies" gorm:"type:jsonb"`                     // Allergies stored as JSONB
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions" gorm:"type:jsonb"`          // Dietary restrictions stored as JSONB
}

// 4. Profile
type Profile struct {
	gorm.Model
	UserID     string `json:"user_id"`                                         // Foreign key
	User       User   `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	Age        int    `json:"age"`                                             // User's age
	Height     int    `json:"height"`                                          // User's height
	Weight     int    `json:"weight"`                                          // User's weight
	HeightUnit string `json:"height_unit"`                                     // Height unit (e.g., cm, inches)
	WeightUnit string `json:"weight_unit"`                                     // Weight unit (e.g., kg, lbs)
	Gender     string `json:"gender"`                                          // Gender of the user
	Lifestyle  string `json:"lifestyle"`                                       // Lifestyle (e.g., sedentary, active)
}

// 5. Diet
// Diet model linked to User
type Diet struct {
	gorm.Model
	UserID          string    `json:"user_id"`                                         // Foreign key
	User            User      `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	PreferredUnits  FoodUnits `json:"preferred_units" gorm:"type:jsonb"`               // Preferred units stored as JSONB
	FavoriteRecipes []string  `json:"favorite_recipes" gorm:"type:jsonb"`              // Favorite recipes stored as JSONB
	DislikedRecipes []string  `json:"disliked_recipes" gorm:"type:jsonb"`              // Disliked recipes stored as JSONB
	FavoriteItems   []string  `json:"favorite_items" gorm:"type:jsonb"`                // Favorite food items stored as JSONB
	DislikedItems   []string  `json:"disliked_items" gorm:"type:jsonb"`                // Disliked food items stored as JSONB
	FavoriteCuisine []string  `json:"favorite_cuisine" gorm:"type:jsonb"`              // Favorite cuisines stored as JSONB
	DislikedCuisine []string  `json:"disliked_cuisine" gorm:"type:jsonb"`              // Disliked cuisines stored as JSONB
}

// 6. Inventory
// Inventory model linked to User
type Inventory struct {
	gorm.Model
	UserID string     `json:"user_id"`                                         // Foreign key
	User   User       `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	Items  []FoodItem `json:"items" gorm:"type:jsonb"`                         // Inventory items stored as JSONB
}

// 7. Shopping
// Shopping model linked to User
type Shopping struct {
	gorm.Model
	UserID                string `json:"user_id"`                                         // Foreign key
	User                  User   `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	WillingnessToShop     bool   `json:"willingness_to_shop"`                             // Boolean for willingness to shop
	BudgetCurrency        string `json:"budget_currency"`                                 // Currency (USD, EUR, etc.)
	Budget                string `json:"budget"`                                          // Budget for shopping
	EaseOfAvailability    string `json:"ease_of_availability"`                            // Ease of availability (hard, medium, easy)
	ShoppingTimeAvailable string `json:"shopping_time"`                                   // Time available for shopping
}

// 8. Goal
// Goal model linked to User
type Goal struct {
	gorm.Model
	UserID     string `json:"user_id"`                                         // Foreign key
	User       User   `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	Goal       string `json:"goal"`                                            // User's goal (e.g., weight loss)
	Target     string `json:"target"`                                          // Target (e.g., 10kg)
	Deadline   string `json:"deadline"`                                        // Deadline for the goal
	Preference string `json:"preference"`                                      // Preference (e.g., low carb, high protein)
	Plan       string `json:"plan"`                                            // Plan to achieve the goal (e.g., diet, exercise)
	Notes      string `json:"notes"`                                           // Notes related to the goal
}

// 9. Script
// Script model linked to User
type Script struct {
	gorm.Model
	UserID string `json:"user_id"`                                         // Foreign key
	User   User   `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	Script string `json:"script"`                                          // Script text for LLM
	Tokens int    `json:"tokens"`                                          // Number of tokens in the script
}

// 10. Authentication
type Authentication struct {
	gorm.Model
	UserID      string `json:"user_id" gorm:"primaryKey"` // User ID
	Passwd      string `json:"passwd"`                    // Password (hashed)
	LastUpdated string `json:"last_updated"`              // Last updated timestamp
}

// 11. LLM
// LLM model linked to User
//type LLM struct {
//	gorm.Model
//	UserID     string `json:"user_id"`                                         // Foreign key
//	User       User   `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
//	LLMChoice  string `json:"llm_choice"`                                      // LLM choice (e.g., GPT-3)
//	LLMVersion string `json:"llm_version"`                                     // Version of the LLM used (e.g., GPT-3.5)
//	LLMToken   string `json:"llm_token"`                                       // API token for LLM usage
//}

// 12. Recipe
// Recipe model optionally linked to User
type Recipe struct {
	gorm.Model
	RecipeID        string               `json:"recipe_id"`                                       // Recipe ID (can be a UUID or string)
	Name            string               `json:"name"`                                            // Recipe name
	Category        string               `json:"category"`                                        // Recipe category (e.g., appetizer, dessert)
	Cuisine         string               `json:"cuisine"`                                         // Cuisine type (e.g., Italian, Mexican)
	Ingredients     []FoodItem           `json:"ingredients" gorm:"type:jsonb"`                   // List of ingredients (stored as JSONB)
	Instructions    []RecipeInstructions `json:"instructions" gorm:"type:jsonb"`                  // Step-by-step instructions (stored as JSONB)
	NutritionalInfo string               `json:"nutritional_info"`                                // Nutritional information
	UserID          *string              `json:"user_id"`                                         // Optional Foreign key to the user who created the recipe (nullable)
	User            *User                `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Optional relationship to User
	RecipeHTML      string               `json:"recipe_html"`                                     // HTML content of the recipe (for rendering purposes)
	Time            RecipeTime           `json:"recipe_time"`                                     // Time required for the recipe (prep time, cooking time)
}

// 13. Choices
// Choices model linked to User
type MealChoice struct {
	gorm.Model
	UserID        string `json:"user_id"`                                         // Foreign key
	User          User   `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	ServingSize   int    `json:"serving_size"`                                    // Serving size preference for recipes
	Shopping      bool   `json:"shopping"`                                        // Whether the user prefers recipes with shopping preferences
	TimeAvailable int    `json:"time_available"`                                  // Time available for cooking
	Innovative    bool   `json:"innovative"`                                      // Whether the user wants to try innovative recipes
	Nutritional   string `json:"nutritional"`                                     // Nutritional preferences (e.g., high protein, low carb)
	MealType      string `json:"meal_type"`                                       // Type of meal (e.g., breakfast, lunch, dinner)
	FineTuned     bool   `json:"fine_tuned"`                                      // Whether fine-tuned recipes are required
}
