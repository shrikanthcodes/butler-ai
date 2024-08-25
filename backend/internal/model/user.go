package model

type UserProfile struct { //struct to store user's profile information (Each user has a profile with basic information like name, email, phone, etc)
	UserID      string `json:"user_id"`      // User ID
	InventoryID string `json:"inventory_id"` // Inventory ID, foreign key
	FirstName   string `json:"first_name"`   //User First Name
	LastName    string `json:"last_name"`    //User Last Name
	Email       string `json:"email"`        //User Email
	Phone       string `json:"phone"`        //User Phone number
	Age         string `json:"age"`          //User Age (25, 30, etc)
	Gender      string `json:"gender"`       //User Gender (Male, Female, Other)
	Weight      string `json:"weight"`       // Weight (120, 150, etc)
	WeightUnit  string `json:"weight_unit"`  // Weight unit (kg, lbs, etc)
	Height      string `json:"height"`       // Height (5'5, 6'0, etc)
	HeightUnit  string `json:"height_unit"`  // Height unit (ft, cm, etc)
}

type UserHealth struct { //struct to store user's health information (Each user has health conditions, medications, allergies, dietary restrictions, etc)
	UserID              string      `json:"user_id"`              // User ID
	HealthConditions    []BasicInfo `json:"health_conditions"`    // Health conditions like asthma, diabetes, etc
	Medications         []BasicInfo `json:"medications"`          // Medications like insulin, etc
	Allergies           []BasicInfo `json:"allergies"`            // Allergies like peanuts, etc
	DietaryRestrictions []BasicInfo `json:"dietary_restrictions"` // Dietary restrictions like vegetarian, vegan, etc
}

type UserAuth struct { //struct to store user's authentication information (Each user has an authentication token, password, last updated timestamp, etc)
	UserID      string `json:"user_id"`      // User ID
	Passwd      string `json:"passwd"`       // Password
	LastUpdated string `json:"last_updated"` // Last updated timestamp
}

type UserDataComplete struct { //struct to store complete user data (User's profile, health information, preferences, inventory, recipe preferences)
	Profile           UserProfile        // User's profile information
	Health            UserHealth         // User's health information
	Preferences       DietaryPreferences // User's preferences for recipes, food items, cuisines, etc
	Inventory         FoodInventory      // User's inventory of food items
	RecipePreferences RecipePreferences  // Recipe preferences for user
}

type UserSession struct { //struct to store user's session refresh token information (Each user has a refresh token, user agent, last login, last logout, etc)
	RefreshToken string `json:"refresh_token"` // Refresh token
	UserID       string `json:"user_id"`       // User ID
	UserAgent    string `json:"user_agent"`    // User agent
	LastLogin    string `json:"last_login"`    // Last login timestamp
	LastLogout   string `json:"last_logout"`   // Last logout timestamp
	ExpiresAt    string `json:"expires_at"`    // Refresh token expiry
}

type UserLLM struct { //struct to store user's LLM information (Each user has an LLM choice, LLM version, LLM token, etc)
	UserID     string `json:"user_id"`     // User ID
	LLMChoice  string `json:"llm_choice"`  // LLM choice
	LLMVersion string `json:"llm_version"` // LLM version
	LLMToken   string `json:"llm_token"`   // LLM token
}

type UserIntegration struct { //struct to store user's integrations information (Each user has integrations with Meta, Google, Twitter, etc)
	UserID           string `json:"user_id"`            // User ID
	Meta             bool   `json:"meta"`               // Boolean to denote whether user has Meta integration
	MetaToken        string `json:"meta_token"`         // Meta token
	MetaExpiresAt    string `json:"meta_expires_at"`    // Meta token expiry
	Google           bool   `json:"google"`             // Boolean to denote whether user has Google integration
	GoogleToken      string `json:"google_token"`       // Google token
	GoogleExpiresAt  string `json:"google_expires_at"`  // Google token expiry
	Twitter          bool   `json:"twitter"`            // Boolean to denote whether user has Twitter integration
	TwitterToken     string `json:"twitter_token"`      // Twitter token
	TwitterExpiresAt string `json:"twitter_expires_at"` // Twitter token expiry
}
