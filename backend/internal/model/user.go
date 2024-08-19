package model

// TODO: Make structs for the following tables

type AuthDB struct {
	UserID      string `json:"user_id"`
	Passwd      string `json:"passwd"`
	LastUpdated string `json:"last_updated"`
}

type Refresh struct {
	RefreshToken string `json:"refresh_token"`
	UserID       string `json:"user_id"`
	UserAgent    string `json:"user_agent"`
	LastLogin    string `json:"last_login"`
	LastLogout   string `json:"last_logout"`
	ExpiresAt    string `json:"expires_at"`
}

type LLM struct {
	UserID     string `json:"user_id"`
	LLMChoice  string `json:"llm_choice"`
	LLMVersion string `json:"llm_version"`
	LLMToken   string `json:"llm_token"`
}

type Integrations struct {
	UserID           string `json:"user_id"`
	Meta             string `json:"meta"`
	MetaToken        string `json:"meta_token"`
	MetaExpiresAt    string `json:"meta_expires_at"`
	Google           string `json:"google"`
	GoogleToken      string `json:"google_token"`
	GoogleExpiresAt  string `json:"google_expires_at"`
	Twitter          string `json:"twitter"`
	TwitterToken     string `json:"twitter_token"`
	TwitterExpiresAt string `json:"twitter_expires_at"`
}

// User DB Tables
type UserProfile struct {
	UserID     string `json:"user_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	Weight     string `json:"weight"`
	WeightUnit string `json:"weight_unit"`
	Height     string `json:"height"`
	HeightUnit string `json:"height_unit"`
}

type Health struct {
	UserID              string      `json:"user_id"`
	HealthConditions    []BasicInfo `json:"health_conditions"`
	Medications         []BasicInfo `json:"medications"`
	Allergies           []BasicInfo `json:"allergies"`
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions"`
}

type Conversations struct {
	ConversationID string     `json:"conversation_id"`
	UserID         string     `json:"user_id"`
	ChatHistory    []Dialogue `json:"chat_history"`
	LastUpdated    string     `json:"last_updated"`
}

type Preferences struct {
	UserID              string      `json:"user_id"`
	FavoriteRecipes     []BasicInfo `json:"favorite_recipes"`
	DislikedRecipes     []BasicInfo `json:"disliked_recipes"`
	FavoriteItems       []BasicInfo `json:"favorite_items"`
	DislikedItems       []BasicInfo `json:"disliked_items"`
	FavoriteCategories  []string    `json:"favorite_categories"`
	DislikedCategories  []string    `json:"disliked_categories"`
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions"`
}

type Inventory struct {
	UserID string       `json:"user_id"`
	Items  []PantryItem `json:"items"`
}

type Recipe struct {
	RecipeID        string               `json:"recipe_id"`
	Name            string               `json:"name"`
	Category        string               `json:"category"`
	Cuisine         string               `json:"cuisine"`
	Ingredients     []PantryItem         `json:"ingredients"`
	Instructions    []RecipeInstructions `json:"instructions"`
	NutritionalInfo string               `json:"nutritional_info"`
	UserID          string               `json:"user_id"`
	RecipeHTML      string               `json:"recipe_html"`
}

type BasicInfo struct {
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
}

type RecipeInstructions struct {
	Intruction_no int    `json:"instruction_no"`
	Instruction   string `json:"instruction"`
	Media         string `json:"media"`
}

// Create a struct for Serving sizes, budget, willingness to shop, etc (user's recipe preferences)
type RecipePreferences struct {
	UserID           string `json:"user_id"`
	ServingSize      string `json:"serving_size"`
	Budget           string `json:"budget"`
	Budget_currency  string `json:"budget_currency"`
	ShopPreference   bool   `json:"shop_preference"`
	ItemAvailability string `json:"item_availability"`
	TimeAvailable    string `json:"time_available"`
	Innovative       bool   `json:"innovative"`
}

type CompleteUserData struct {
	Profile           UserProfile
	Health            Health
	Preferences       Preferences
	Inventory         Inventory
	RecipePreferences RecipePreferences
}
