package user

import (
	config "backend/internal/config"
)

func GetUserProfile(userID string) config.UserProfile {
	// Try to use userID to get user profile or Create fake data (test)
	profile := config.UserProfile{
		UserID:     "123",
		FirstName:  "Aishwarya",
		LastName:   "Srinivasan",
		Email:      "aishwarya@gmail.com",
		Phone:      "1234567890",
		Age:        "25",
		Gender:     "Female",
		Weight:     "120",
		WeightUnit: "lbs",
		Height:     "5'5",
		HeightUnit: "ft",
	}
	return profile
}

func GetUserHealth(userID string) config.Health {
	// Try to use userID to get user health or Create fake data (test)
	health := config.Health{
		UserID: "123",
		HealthConditions: []config.BasicInfo{
			{ItemName: "Asthma", Description: "Mild case of asthma"},
			{ItemName: "Diabetes", Description: "Type 1 diabetes, controlled with insulin"},
		},
		Medications: []config.BasicInfo{
			{ItemName: "Insulin", Description: "10 units before meals"},
		},
		Allergies: []config.BasicInfo{
			{ItemName: "Peanuts", Description: "Mild allergy to peanuts, especially when eaten raw or at night"},
		},
		DietaryRestrictions: []config.BasicInfo{
			{ItemName: "Vegetarian", Description: "Vegetarian diet, no meat or fish. Eggs and dairy are okay"},
		},
	}
	return health
}

func GetUserPreferences(userID string) config.Preferences {
	// Try to use userID to get user preferences or Create fake data (test)
	preferences := config.Preferences{
		UserID: "123",
		FavoriteRecipes: []config.BasicInfo{
			{ItemName: "Pasta", Description: "All types of pasta, especially with tomato sauce"},
		},
		DislikedRecipes: []config.BasicInfo{
			{ItemName: "Fish curry", Description: "Any type of fish, especially salmon"},
		},
		FavoriteItems: []config.BasicInfo{
			{ItemName: "Tomatoes", Description: "All types of tomatoes, especially cherry tomatoes"},
		},
		DislikedItems: []config.BasicInfo{
			{ItemName: "Mushrooms", Description: "All types of mushrooms, especially shiitake"},
		},
		FavoriteCategories: []string{
			"Indian",
			"Mexican",
		},
		DislikedCategories: []string{
			"Chinese",
			"Japanese",
		},
		DietaryRestrictions: []config.BasicInfo{
			{ItemName: "Vegetarian", Description: "No meat or fish. Eggs and dairy are okay"},
		},
	}
	return preferences
}

func GetUserInventory(userID string) config.Inventory {
	// Try to use userID to get user inventory or Create fake data (test)
	inventory := config.Inventory{
		UserID: "123",
		Items: []config.PantryItem{
			{ItemName: "Tomatoes", Quantity: 5, Unit: "pieces"},
			{ItemName: "Onions", Quantity: 3, Unit: "pieces"},
			{ItemName: "Garlic", Quantity: 2, Unit: "cloves"},
		},
	}
	return inventory
}

func GetUserRecipePreferences(userID string) config.RecipePreferences {
	// Try to use userID to get user recipe preferences or Create fake data (test)
	recipePreferences := config.RecipePreferences{
		UserID:           "123",
		ServingSize:      "2",
		Budget:           "Medium",
		ShopPreference:   true,
		ItemAvailability: "Very Common",
	}
	return recipePreferences
}

// GetCompleteUserData aggregates all user-related data.
func GetCompleteUserData(userID string) config.CompleteUserData {
	profile := GetUserProfile(userID)
	health := GetUserHealth(userID)
	preferences := GetUserPreferences(userID)
	inventory := GetUserInventory(userID)
	recipePreferences := GetUserRecipePreferences(userID)

	return config.CompleteUserData{
		Profile:           profile,
		Health:            health,
		Preferences:       preferences,
		Inventory:         inventory,
		RecipePreferences: recipePreferences,
	}
}
