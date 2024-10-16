package entity

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time ` json:"deleted_at"`
}

// User defines the structure of the basic user table
type User struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	BaseModel
}

// Writeup defines the structure of the writeup table
type Writeup struct {
	UserID string `json:"user_id"`
	BaseModel

	RecipeWriteup         string `json:"recipe_writeup"`
	ShoppingWriteup       string `json:"shopping_writeup"`
	CalorieTrackerWriteup string `json:"calorie_tracker_writeup"`
	HealthWriteup         string `json:"health_writeup"`
	MotivationWriteup     string `json:"motivation_writeup"`
}

// Conversation defines the structure of the conversation table
type Conversation struct {
	ConvID      string            `json:"conv_id"`
	UserID      string            `json:"user_id"`
	Title       string            `json:"title"`
	Transcript  JSONB[[]Dialogue] `json:"conversation"`
	LastUpdated *time.Time        `json:"last_updated"`
	Summary     *string           `json:"summary"`
	ChatType    string            `json:"chat_type"`
	IsActive    bool              `json:"is_active"`
	BaseModel
}

// Medical defines the structure of the health table
type Medical struct {
	UserID              string             `json:"user_id"`
	HealthConditions    JSONB[[]BasicInfo] `json:"health_conditions"`
	Medications         JSONB[[]BasicInfo] `json:"medications"`
	Allergies           JSONB[[]BasicInfo] `json:"allergies"`
	DietaryRestrictions JSONB[[]BasicInfo] `json:"dietary_restrictions"`
	BaseModel
}

// Profile defines the structure of the profile table
type Profile struct {
	UserID     string `json:"user_id"`
	Age        int    `json:"age"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	HeightUnit string `json:"height_unit"`
	WeightUnit string `json:"weight_unit"`
	Gender     string `json:"gender"`
	Lifestyle  string `json:"lifestyle"`
	BaseModel
}

// Diet defines the structure of the diet table
type Diet struct {
	UserID          string          `json:"user_id"`
	PreferredUnits  string          `json:"preferred_units"`
	FavoriteRecipes JSONB[[]string] `json:"favorite_recipes"`
	DislikedRecipes JSONB[[]string] `json:"disliked_recipes"`
	FavoriteItems   JSONB[[]string] `json:"favorite_items"`
	DislikedItems   JSONB[[]string] `json:"disliked_items"`
	FavoriteCuisine JSONB[[]string] `json:"favorite_cuisine"`
	DislikedCuisine JSONB[[]string] `json:"disliked_cuisine"`
	BaseModel
}

// Inventory defines the structure of the inventory table
type Inventory struct {
	UserID string            `json:"user_id"`
	Items  JSONB[[]FoodItem] `json:"items"`
	BaseModel
}

// Shopping defines the structure of the shopping table
type Shopping struct {
	UserID             string            `json:"user_id"`
	ShoppingMode       string            `json:"shopping_mode"`
	BudgetCurrency     string            `json:"budget_currency"`
	Budget             int               `json:"budget"`
	EaseOfAvailability string            `json:"ease_of_availability"`
	ShoppingList       JSONB[[]FoodItem] `json:"shopping_list"`
	BaseModel
}

// Goal defines the structure of the goal table
type Goal struct {
	GoalID     string          `json:"goal_id"` // Previously gorm:"primaryKey"`
	UserID     string          `json:"user_id"`
	Goal       string          `json:"goal"`
	Target     string          `json:"target"`
	Deadline   string          `json:"deadline"`
	Preference string          `json:"preference"`
	Plan       JSONB[[]string] `json:"plan"`
	Notes      string          `json:"notes"`
	BaseModel
}

// Authentication defines the structure of the authentication table
type Authentication struct {
	UserID      string     `json:"user_id"`
	Role        string     `json:"role"`
	Passwd      string     `json:"passwd"`
	LastUpdated *time.Time `json:"last_updated"`
	AuthToken   string     `json:"auth_token"`
	BaseModel
}

// Recipe defines the structure of the recipe table
type Recipe struct {
	RecipeID        string                      `json:"recipe_id"`
	Name            string                      `json:"name"`
	Tags            JSONB[[]string]             `json:"tags"`
	Cuisine         string                      `json:"cuisine"`
	Ingredients     JSONB[[]FoodItem]           `json:"ingredients"`
	Instructions    JSONB[[]RecipeInstructions] `json:"instructions"`
	NutritionalInfo JSONB[[]FoodItem]           `json:"nutritional_info"`
	UserID          *string                     `json:"user_id"`
	URL             string                      `json:"url"`
	Time            string                      `json:"recipe_time"`
	BaseModel
}

// MealChoice defines the structure of the meal choice table
type MealChoice struct {
	UserID        string          `json:"user_id"`
	ServingSize   int             `json:"serving_size"`
	Shopping      bool            `json:"shopping"`
	TimeAvailable int             `json:"time_available"`
	Innovative    bool            `json:"innovative"`
	NutritionTag  JSONB[[]string] `json:"nutrition_tag"`
	MealType      JSONB[[]string] `json:"meal_type"`
	FineTuned     bool            `json:"fine_tuned"`
	BaseModel
}
