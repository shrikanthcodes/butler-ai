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

func GetUserHealth(userID string) model.UserHealth {
	// Try to use userID to get user health or Create fake data (test)
	health := model.UserHealth{
		UserID: "123",
		HealthConditions: []model.BasicInfo{
			{Name: "Asthma", Description: "Mild case of asthma"},
			{Name: "Diabetes", Description: "Type 1 diabetes, controlled with insulin"},
		},
		Medications: []model.BasicInfo{
			{Name: "Insulin", Description: "10 units before meals"},
		},
		Allergies: []model.BasicInfo{
			{Name: "Peanuts", Description: "Mild allergy to peanuts, especially when eaten raw or at night"},
		},
		DietaryRestrictions: []model.BasicInfo{
			{Name: "Vegetarian", Description: "Vegetarian diet, no meat or fish. Eggs and dairy are okay"},
		},
	}
	return health
}

func GetUserPreferences(userID string) model.DietaryPreferences {
	// Try to use userID to get user preferences or Create fake data (test)
	preferences := model.DietaryPreferences{
		UserID:          "123",
		FavoriteRecipes: []string{"Pasta Bolognese", "Chicken curry"},
		DislikedRecipes: []string{"Mushroom risotto", "Sushi"},
		FavoriteItems:   []string{"Tomatoes", "Onions", "Shrimp"},
		DislikedItems:   []string{"Mushrooms", "Tofu"},
		FavoriteCuisine: []string{"Indian", "Mexican"},
		DislikedCuisine: []string{"Chinese", "Japanese"},
	}
	return preferences
}

func GetUserInventory(userID string) model.FoodInventory {
	// Try to use userID to get user inventory or Create fake data (test)
	inventory := model.FoodInventory{
		UserID: "123",
		Items: []model.FoodItem{
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
		UserID:      "123",
		ServingSize: 2,
	}
	return recipePreferences
}

// GetCompleteUserData aggregates all user-related data.
func GetCompleteUserData(userID string) model.UserDataComplete {
	profile := GetUserProfile(userID)
	health := GetUserHealth(userID)
	preferences := GetUserPreferences(userID)
	inventory := GetUserInventory(userID)
	recipePreferences := GetUserRecipePreferences(userID)

	return model.UserDataComplete{
		Profile:           profile,
		Health:            health,
		Preferences:       preferences,
		Inventory:         inventory,
		RecipePreferences: recipePreferences,
	}
}
