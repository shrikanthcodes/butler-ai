package models

import (
	enum "backend/pkg/models/enum"

	gorm "gorm.io/gorm"
)

// 1. User
type User struct {
	gorm.Model
	UserID string `json:"user_id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Phone  string `json:"phone" gorm:"unique" optional:"true"`

	// Relationships to other tables
	Auth         *Authentication `gorm:"foreignKey:UserID"` // Optional relationship
	Writeup      *Writeup        `gorm:"foreignKey:UserID"` // Optional relationship
	Conversation []Conversation  `gorm:"foreignKey:UserID"` // Optional relationship
	Profile      *Profile        `gorm:"foreignKey:UserID"` // Optional relationship
	Health       *Health         `gorm:"foreignKey:UserID"` // Optional relationship
	Diet         *Diet           `gorm:"foreignKey:UserID"` // Optional relationship
	Inventory    *Inventory      `gorm:"foreignKey:UserID"` // Optional relationship
	Goal         *Goal           `gorm:"foreignKey:UserID"` // Optional relationship
	Recipe       *Recipe         `gorm:"foreignKey:UserID"` // Optional relationship
	Shopping     *Shopping       `gorm:"foreignKey:UserID"` // Optional relationship
	MealChoices  *MealChoice     `gorm:"foreignKey:UserID"` // Optional relationship
}

type Writeup struct {
	gorm.Model
	UserID string `json:"user_id" gorm:"primaryKey"` // Foreign key
	// Writeups for each of the chat types
	RecipeWriteup         string `json:"recipe_writeup"`          // Writeup for recipe chat
	ShoppingWriteup       string `json:"shopping_writeup"`        // Writeup for shopping chat
	CalorieTrackerWriteup string `json:"calorie_tracker_writeup"` // Writeup for calorie tracker chat
	HealthWriteup         string `json:"health_writeup"`          // Writeup for health chat
	MotivationWriteup     string `json:"motivation_writeup"`      // Writeup for motivation chat
}

// 2. Conversation
type Conversation struct {
	gorm.Model
	ConvID          string        `json:"conv_id" gorm:"primaryKey"`                       // Primary key
	UserID          string        `json:"user_id"`                                         // Foreign key
	User            User          `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Define foreign key
	Conversation    []Dialogue    `json:"conversation" gorm:"type:jsonb"`                  // Actual conversation between user and AI
	LastUpdated     string        `json:"last_updated" optional:"true"`                    // Last updated timestamp
	Summary         string        `json:"summary" optional:"true"`                         // Summary of the conversation
	RecentDialogues []Dialogue    `json:"recent_dialogues" gorm:"type:jsonb"`              // Recent dialogues stored as JSONB
	ChatType        enum.ChatType `json:"chat_type"`                                       // Type of chat (e.g., recipe, shopping)
	IsActive        bool          `json:"is_active"`
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
	UserID     string          `json:"user_id"`                                         // Foreign key
	User       User            `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	Age        int             `json:"age"`                                             // User's age
	Height     int             `json:"height"`                                          // User's height
	Weight     int             `json:"weight"`                                          // User's weight
	HeightUnit enum.HeightUnit `json:"height_unit"`                                     // Height unit (enum)
	WeightUnit enum.WeightUnit `json:"weight_unit"`                                     // Weight unit (enum)
	Gender     enum.Gender     `json:"gender"`                                          // Gender of the user
	Lifestyle  enum.Lifestyle  `json:"lifestyle"`                                       // Lifestyle (enum)
}

// 5. Diet
// Diet model linked to User
type Diet struct {
	gorm.Model
	UserID          string          `json:"user_id"`                                         // Foreign key
	User            User            `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	PreferredUnits  FoodUnits       `json:"preferred_units" gorm:"type:jsonb"`               // Preferred units stored as JSONB
	FavoriteRecipes []string        `json:"favorite_recipes" gorm:"type:jsonb"`              // Favorite recipes stored as JSONB
	DislikedRecipes []string        `json:"disliked_recipes" gorm:"type:jsonb"`              // Disliked recipes stored as JSONB
	FavoriteItems   []string        `json:"favorite_items" gorm:"type:jsonb"`                // Favorite food items stored as JSONB
	DislikedItems   []string        `json:"disliked_items" gorm:"type:jsonb"`                // Disliked food items stored as JSONB
	FavoriteCuisine []enum.Cuisines `json:"favorite_cuisine" gorm:"type:jsonb"`              // Favorite cuisines stored as JSONB
	DislikedCuisine []enum.Cuisines `json:"disliked_cuisine" gorm:"type:jsonb"`              // Disliked cuisines stored as JSONB
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
	UserID             string          `json:"user_id"`                                         // Foreign key
	User               User            `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	ShoppingMode       enum.ShopMode   `json:"shopping_mode"`                                   // Boolean for willingness to shop
	BudgetCurrency     enum.Currencies `json:"budget_currency"`                                 // Currency (USD, EUR, etc.)
	Budget             int             `json:"budget"`                                          // Budget for shopping
	EaseOfAvailability enum.Difficulty `json:"ease_of_availability"`                            // Ease of availability (hard, medium, easy)
	ShoppingList       []FoodItem      `json:"shopping_list" gorm:"type:jsonb"`                 // Shopping list stored as JSONB
}

// 8. Goal
// Goal model linked to User
type Goal struct {
	gorm.Model
	GoalID     string         `json:"goal_id" gorm:"primaryKey"`                       // Goal ID
	UserID     string         `json:"user_id"`                                         // Foreign key
	User       User           `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Foreign key relationship to User
	Goal       enum.GoalTypes `json:"goal"`                                            // User's goal (e.g., weight loss)
	Target     string         `json:"target"`                                          // Target (e.g., 10kg)
	Deadline   string         `json:"deadline"`                                        // Deadline (datetime) for the goal
	Preference string         `json:"preference"`                                      // Preference (e.g., low carb, high protein)
	Plan       []string       `json:"plan" gorm:"type:jsonb"`                          // Plan to achieve the goal (step by step)
	Notes      string         `json:"notes"`                                           // Notes related to the goal
}

// 10. Authentication
type Authentication struct {
	gorm.Model
	UserID      string `json:"user_id" gorm:"primaryKey"` // User ID
	Passwd      string `json:"passwd"`                    // Password (hashed)
	LastUpdated string `json:"last_updated"`              // Last updated timestamp
}

// 12. Recipe
// Recipe model optionally linked to User
type Recipe struct {
	gorm.Model
	RecipeID        string               `json:"recipe_id"`                                       // Recipe ID (can be a UUID or string)
	Name            string               `json:"name"`                                            // Recipe name
	Tags            []enum.RecipeTags    `json:"tags" gorm:"type:jsonb"`                          // Recipe tags (e.g., appetizer, dessert)
	Cuisine         enum.Cuisines        `json:"cuisine"`                                         // Cuisine type (e.g., Italian, Mexican)
	Ingredients     []FoodItem           `json:"ingredients" gorm:"type:jsonb"`                   // List of ingredients (stored as JSONB)
	Instructions    []RecipeInstructions `json:"instructions" gorm:"type:jsonb"`                  // Step-by-step instructions (stored as JSONB)
	NutritionalInfo []FoodItem           `json:"nutritional_info" gorm:"type:jsonb"`              // Nutritional information
	UserID          *string              `json:"user_id"`                                         // Optional Foreign key to the user who created the recipe (nullable)
	User            *User                `json:"user" gorm:"foreignKey:UserID;references:UserID"` // Optional relationship to User
	URL             string               `json:"url"`                                             // HTML content of the recipe (for rendering purposes)
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
	FineTuned     bool   `json:"fine_tuned" default:"false"`                      // Whether fine-tuned recipes are required
}
