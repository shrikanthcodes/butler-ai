package models

import (
	"time"

	enum "github.com/shrikanthcodes/butler-ai/backend/pkg/models/enum"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// 1. User
type User struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Phone  string `json:"phone" gorm:"unique"`
	BaseModel
}

// 2. Writeup
type Writeup struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	User   User   `gorm:"foreignKey:UserID;references:UserID"`
	BaseModel

	RecipeWriteup         string `json:"recipe_writeup"`
	ShoppingWriteup       string `json:"shopping_writeup"`
	CalorieTrackerWriteup string `json:"calorie_tracker_writeup"`
	HealthWriteup         string `json:"health_writeup"`
	MotivationWriteup     string `json:"motivation_writeup"`
}

// 3. Conversation
type Conversation struct {
	ConvID          string        `json:"conv_id" gorm:"primaryKey"`
	UserID          string        `json:"user_id"`
	User            User          `gorm:"foreignKey:UserID;references:UserID"`
	Conversation    DialogueArray `json:"conversation" gorm:"type:jsonb"`
	LastUpdated     *time.Time    `json:"last_updated"`
	Summary         *string       `json:"summary"`
	RecentDialogues DialogueArray `json:"recent_dialogues" gorm:"type:jsonb"`
	ChatType        enum.ChatType `json:"chat_type"`
	IsActive        bool          `json:"is_active"`
	BaseModel
}

// 4. Health
type Health struct {
	UserID              string      `json:"user_id" gorm:"primaryKey"`
	User                User        `gorm:"foreignKey:UserID;references:UserID"`
	HealthConditions    []BasicInfo `json:"health_conditions" gorm:"type:jsonb"`
	Medications         []BasicInfo `json:"medications" gorm:"type:jsonb"`
	Allergies           []BasicInfo `json:"allergies" gorm:"type:jsonb"`
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions" gorm:"type:jsonb"`
	BaseModel
}

// 5. Profile
type Profile struct {
	UserID     string          `json:"user_id" gorm:"primaryKey"`
	User       User            `gorm:"foreignKey:UserID;references:UserID"`
	Age        int             `json:"age"`
	Height     int             `json:"height"`
	Weight     int             `json:"weight"`
	HeightUnit enum.HeightUnit `json:"height_unit"`
	WeightUnit enum.WeightUnit `json:"weight_unit"`
	Gender     enum.Gender     `json:"gender"`
	Lifestyle  enum.Lifestyle  `json:"lifestyle"`
	BaseModel
}

// 6. Diet
type Diet struct {
	UserID          string          `json:"user_id" gorm:"primaryKey"`
	User            User            `gorm:"foreignKey:UserID;references:UserID"`
	PreferredUnits  FoodUnits       `json:"preferred_units" gorm:"type:jsonb"`
	FavoriteRecipes []string        `json:"favorite_recipes" gorm:"type:jsonb"`
	DislikedRecipes []string        `json:"disliked_recipes" gorm:"type:jsonb"`
	FavoriteItems   []string        `json:"favorite_items" gorm:"type:jsonb"`
	DislikedItems   []string        `json:"disliked_items" gorm:"type:jsonb"`
	FavoriteCuisine []enum.Cuisines `json:"favorite_cuisine" gorm:"type:jsonb"`
	DislikedCuisine []enum.Cuisines `json:"disliked_cuisine" gorm:"type:jsonb"`
	BaseModel
}

// 7. Inventory
type Inventory struct {
	UserID string     `json:"user_id" gorm:"primaryKey"`
	User   User       `gorm:"foreignKey:UserID;references:UserID"`
	Items  []FoodItem `json:"items" gorm:"type:jsonb"`
	BaseModel
}

// 8. Shopping
type Shopping struct {
	UserID             string          `json:"user_id" gorm:"primaryKey"`
	User               User            `gorm:"foreignKey:UserID;references:UserID"`
	ShoppingMode       enum.ShopMode   `json:"shopping_mode"`
	BudgetCurrency     enum.Currencies `json:"budget_currency"`
	Budget             int             `json:"budget"`
	EaseOfAvailability enum.Difficulty `json:"ease_of_availability"`
	ShoppingList       []FoodItem      `json:"shopping_list" gorm:"type:jsonb"`
	BaseModel
}

// 9. Goal
type Goal struct {
	GoalID     string         `json:"goal_id" gorm:"primaryKey"`
	UserID     string         `json:"user_id"`
	User       User           `gorm:"foreignKey:UserID;references:UserID"`
	Goal       enum.GoalTypes `json:"goal"`
	Target     string         `json:"target"`
	Deadline   string         `json:"deadline"`
	Preference string         `json:"preference"`
	Plan       []string       `json:"plan" gorm:"type:jsonb"`
	Notes      string         `json:"notes"`
	BaseModel
}

// 10. Authentication
type Authentication struct {
	UserID      string     `json:"user_id" gorm:"primaryKey"`
	User        User       `gorm:"foreignKey:UserID;references:UserID"`
	Passwd      string     `json:"passwd"`
	LastUpdated *time.Time `json:"last_updated"`
	BaseModel
}

// 11. Recipe
type Recipe struct {
	RecipeID        string               `json:"recipe_id" gorm:"primaryKey"`
	Name            string               `json:"name"`
	Tags            []enum.RecipeTags    `json:"tags" gorm:"type:jsonb"`
	Cuisine         enum.Cuisines        `json:"cuisine"`
	Ingredients     []FoodItem           `json:"ingredients" gorm:"type:jsonb"`
	Instructions    []RecipeInstructions `json:"instructions" gorm:"type:jsonb"`
	NutritionalInfo []FoodItem           `json:"nutritional_info" gorm:"type:jsonb"`
	UserID          *string              `json:"user_id"`
	User            *User                `gorm:"foreignKey:UserID;references:UserID"`
	URL             string               `json:"url"`
	Time            RecipeTime           `json:"recipe_time" gorm:"type:jsonb"`
	BaseModel
}

// 12. MealChoice
type MealChoice struct {
	UserID        string `json:"user_id" gorm:"primaryKey"`
	User          User   `gorm:"foreignKey:UserID;references:UserID"`
	ServingSize   int    `json:"serving_size"`
	Shopping      bool   `json:"shopping"`
	TimeAvailable int    `json:"time_available"` //Add units (hardcode to mins)
	Innovative    bool   `json:"innovative"`
	Nutritional   string `json:"nutritional"` //convert to fooditem
	MealType      string `json:"meal_type"`   //enum
	FineTuned     bool   `json:"fine_tuned" gorm:"default:false"`
	BaseModel
}
