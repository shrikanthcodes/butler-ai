package user

import (
	model "backend/internal/model"
)

func GetUserProfile(userID string) model.UserProfile {
	// Try to use userID to get user profile or Create fake data (test)
	profile := model.UserProfile{
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

func GetUserHealth(userID string) model.Health {
	// Try to use userID to get user health or Create fake data (test)
	health := model.Health{
		UserID: "123",
		HealthConditions: []model.BasicInfo{
			{ItemName: "Asthma", Description: "Mild case of asthma"},
			{ItemName: "Diabetes", Description: "Type 1 diabetes, controlled with insulin"},
		},
		Medications: []model.BasicInfo{
			{ItemName: "Insulin", Description: "10 units before meals"},
		},
		Allergies: []model.BasicInfo{
			{ItemName: "Peanuts", Description: "Mild allergy to peanuts, especially when eaten raw or at night"},
		},
		DietaryRestrictions: []model.BasicInfo{
			{ItemName: "Vegetarian", Description: "Vegetarian diet, no meat or fish. Eggs and dairy are okay"},
		},
	}
	return health
}

func GetUserPreferences(userID string) model.Preferences {
	// Try to use userID to get user preferences or Create fake data (test)
	preferences := model.Preferences{
		UserID: "123",
		FavoriteRecipes: []model.BasicInfo{
			{ItemName: "Pasta", Description: "All types of pasta, especially with tomato sauce"},
		},
		DislikedRecipes: []model.BasicInfo{
			{ItemName: "Fish curry", Description: "Any type of fish, especially salmon"},
		},
		FavoriteItems: []model.BasicInfo{
			{ItemName: "Tomatoes", Description: "All types of tomatoes, especially cherry tomatoes"},
		},
		DislikedItems: []model.BasicInfo{
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
		DietaryRestrictions: []model.BasicInfo{
			{ItemName: "Vegetarian", Description: "No meat or fish. Eggs and dairy are okay"},
		},
	}
	return preferences
}

func GetUserInventory(userID string) model.Inventory {
	// Try to use userID to get user inventory or Create fake data (test)
	inventory := model.Inventory{
		UserID: "123",
		Items: []model.PantryItem{
			{ItemName: "Tomatoes", Quantity: 5, Unit: "pieces"},
			{ItemName: "Onions", Quantity: 3, Unit: "pieces"},
			{ItemName: "Garlic", Quantity: 2, Unit: "cloves"},
		},
	}
	return inventory
}

func GetUserRecipePreferences(userID string) model.RecipePreferences {
	// Try to use userID to get user recipe preferences or Create fake data (test)
	recipePreferences := model.RecipePreferences{
		UserID:           "123",
		ServingSize:      "2",
		Budget:           "Medium",
		ShopPreference:   true,
		ItemAvailability: "Very Common",
	}
	return recipePreferences
}

// GetCompleteUserData aggregates all user-related data.
func GetCompleteUserData(userID string) model.CompleteUserData {
	profile := GetUserProfile(userID)
	health := GetUserHealth(userID)
	preferences := GetUserPreferences(userID)
	inventory := GetUserInventory(userID)
	recipePreferences := GetUserRecipePreferences(userID)

	return model.CompleteUserData{
		Profile:           profile,
		Health:            health,
		Preferences:       preferences,
		Inventory:         inventory,
		RecipePreferences: recipePreferences,
	}
}
