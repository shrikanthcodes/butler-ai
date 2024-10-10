package entity

import (
	"github.com/shrikanthcodes/butler-ai/internal/entity/enum"
	"google.golang.org/api/mirror/v1"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// User defines the structure of the basic user table
type User struct {
	UserID string `json:"user_id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Phone  string `json:"phone" gorm:"unique"`
	BaseModel
}

// Writeup defines the structure of the writeup table
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

// Conversation defines the structure of the conversation table
type Conversation struct {
	ConvID          string         `json:"conv_id" gorm:"primaryKey"`
	UserID          string         `json:"user_id"`
	User            User           `gorm:"foreignKey:UserID;references:UserID"`
	Conversation    DialogueArray  `json:"conversation" gorm:"type:jsonb"`
	LastUpdated     *time.Time     `json:"last_updated"`
	Summary         *string        `json:"summary"`
	RecentDialogues DialogueArray  `json:"recent_dialogues" gorm:"type:jsonb"`
	ChatType        enum.ChatTypes `json:"chat_type"`
	IsActive        bool           `json:"is_active"`
	BaseModel
}

// Health defines the structure of the health table
type Health struct {
	UserID              string      `json:"user_id" gorm:"primaryKey"`
	User                User        `gorm:"foreignKey:UserID;references:UserID"`
	HealthConditions    []BasicInfo `json:"health_conditions" gorm:"type:jsonb"`
	Medications         []BasicInfo `json:"medications" gorm:"type:jsonb"`
	Allergies           []BasicInfo `json:"allergies" gorm:"type:jsonb"`
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions" gorm:"type:jsonb"`
	BaseModel
}

// Profile defines the structure of the profile table
type Profile struct {
	UserID     string           `json:"user_id" gorm:"primaryKey"`
	User       User             `gorm:"foreignKey:UserID;references:UserID"`
	Age        int              `json:"age"`
	Height     int              `json:"height"`
	Weight     int              `json:"weight"`
	HeightUnit enum.HeightUnits `json:"height_unit"`
	WeightUnit enum.WeightUnits `json:"weight_unit"`
	Gender     enum.Genders     `json:"gender"`
	Lifestyle  enum.Lifestyles  `json:"lifestyle"`
	BaseModel
}

// Diet defines the structure of the diet table
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

// Inventory defines the structure of the inventory table
type Inventory struct {
	UserID string     `json:"user_id" gorm:"primaryKey"`
	User   User       `gorm:"foreignKey:UserID;references:UserID"`
	Items  []FoodItem `json:"items" gorm:"type:jsonb"`
	BaseModel
}

// Shopping defines the structure of the shopping table
type Shopping struct {
	UserID             string             `json:"user_id" gorm:"primaryKey"`
	User               User               `gorm:"foreignKey:UserID;references:UserID"`
	ShoppingMode       enum.ShoppingTypes `json:"shopping_mode"`
	BudgetCurrency     enum.Currencies    `json:"budget_currency"`
	Budget             int                `json:"budget"`
	EaseOfAvailability enum.Difficulties  `json:"ease_of_availability"`
	ShoppingList       []FoodItem         `json:"shopping_list" gorm:"type:jsonb"`
	BaseModel
}

// Goal defines the structure of the goal table
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

// Authentication defines the structure of the authentication table
type Authentication struct {
	UserID           string     `json:"user_id" gorm:"primaryKey"`
	User             User       `gorm:"foreignKey:UserID;references:UserID"`
	Role             enum.Roles `json:"role"`
	Passwd           string     `json:"passwd"`
	LastUpdated      *time.Time `json:"last_updated"`
	mirror.AuthToken            // Embedded struct
	BaseModel
}

// Recipe defines the structure of the recipe table
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

// MealChoice defines the structure of the meal choice table
type MealChoice struct {
	UserID        string               `json:"user_id" gorm:"primaryKey"`
	User          User                 `gorm:"foreignKey:UserID;references:UserID"`
	ServingSize   int                  `json:"serving_size"`
	Shopping      bool                 `json:"shopping"`
	TimeAvailable int                  `json:"time_available"`
	Innovative    bool                 `json:"innovative"`
	NutritionTag  []enum.NutritionTags `json:"nutrition_tag" gorm:"type:jsonb"`
	MealType      []enum.MealTypes     `json:"meal_type" gorm:"type:jsonb"`
	FineTuned     bool                 `json:"fine_tuned" gorm:"default:false"`
	BaseModel
}
