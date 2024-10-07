package enum

func (g Gender) String() string {
	if g < Unknown || g > NonBinary {
		return "Unknown"
	}
	return [...]string{"Unknown", "Male", "Female", "Non-Binary"}[g]
}

func (h HeightUnit) String() string {
	if h < Centimeters || h > Inches {
		return "Inches"
	}
	return [...]string{"Centimeters", "Inches"}[h]
}

func (w WeightUnit) String() string {
	if w < Kilograms || w > Pounds {
		return "Pounds"
	}
	return [...]string{"Kilograms", "Pounds"}[w]
}

func (l Lifestyle) String() string {
	if l < Sedentary || l > Active {
		return "Medium"
	}
	return [...]string{"Sedentary", "Average", "Active"}[l]
}

func (c ChatType) String() string {
	if c < RecipeChat || c > CalorieTrackerChat {
		return "Unknown"
	}
	return [...]string{"Recipe", "Shopping", "Health", "Motivation", "Calorie Tracker"}[c]
}

func (c Cuisines) String() string {
	if c < African || c > Asian {
		return "Unknown"
	}
	return [...]string{"African", "American", "Chinese", "French", "Indian", "Italian", "Japanese", "Mexican", "Thai", "Filipino", "Asian"}[c]
}

func (c Currencies) String() string {
	if c < USD || c > PHP {
		return "USD"
	}
	return [...]string{"USD", "EUR", "GBP", "JPY", "CNY", "INR", "MXN", "THB", "PHP"}[c]
}

func (t TimeUnits) String() string {
	if t < Seconds || t > Weeks {
		return "Unknown"
	}
	return [...]string{"Seconds", "Minutes", "Hours", "Days", "Weeks"}[t]
}

func (r RecipeTags) String() string {
	if r < MainCourse || r > Snack {
		return "Unknown"
	}
	return [...]string{"Main Course", "Side Dish", "Dessert", "Appetizer", "Salad", "Bread", "Breakfast", "Soup", "Beverage", "Sauce", "Marinade", "Preserves", "Spice Mix", "Dinner", "Lunch", "Snack"}[r]
}

func (d Difficulty) String() string {
	if d < Easy || d > Hard {
		return "Unknown"
	}
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

func (s ShopMode) String() string {
	if s < Online || s > Pickup {
		return "Unknown"
	}
	return [...]string{"Online", "In-person", "Delivery", "Pickup"}[s]
}

func (g GoalTypes) String() string {
	if g < WeightLoss || g > Discipline {
		return "Unknown"
	}
	return [...]string{"Weight Loss", "Weight Gain", "Self Love", "Addiction", "Custom", "New Skill", "Discipline"}[g]
}
