package enum

type Genders int

const (
	Gender0 Genders = iota
	Gender1
	Gender2
	Gender3
)

func (g Genders) String() string {
	if g < Gender0 || g > Gender3 {
		return "Unknown"
	}
	return [...]string{"Unknown", "Male", "Female", "Non-Binary"}[g]
}

type HeightUnits int

const (
	HeightUnit0 HeightUnits = iota
	HeightUnit1
)

func (h HeightUnits) String() string {
	if h < HeightUnit0 || h > HeightUnit1 {
		return "Inches"
	}
	return [...]string{"Centimeters", "Inches"}[h]
}

type WeightUnits int

const (
	WeightUnit0 WeightUnits = iota
	WeightUnit1
)

func (w WeightUnits) String() string {
	if w < WeightUnit0 || w > WeightUnit1 {
		return "Pounds"
	}
	return [...]string{"Kilograms", "Pounds"}[w]
}

type Lifestyles int

const (
	Lifestyle0 Lifestyles = iota
	Lifestyle1
	Lifestyle2
)

func (l Lifestyles) String() string {
	if l < Lifestyle0 || l > Lifestyle2 {
		return "Medium"
	}
	return [...]string{"Sedentary", "Average", "Active"}[l]
}

type ChatTypes int

const (
	ChatType0 ChatTypes = iota
	ChatType1
	ChatType2
	ChatType3
	ChatType4
)

func (c ChatTypes) String() string {
	if c < ChatType0 || c > ChatType4 {
		return "Unknown"
	}
	return [...]string{"Recipe", "Shopping", "Health", "Motivation", "Calorie Tracker"}[c]
}

type Cuisines int

const (
	Cuisine0 Cuisines = iota
	Cuisine1
	Cuisine2
	Cuisine3
	Cuisine4
	Cuisine5
	Cuisine6
	Cuisine7
	Cuisine8
	Cuisine9
	Cuisine10
)

func (c Cuisines) String() string {
	if c < Cuisine0 || c > Cuisine10 {
		return "Unknown"
	}
	return [...]string{"African", "American", "Chinese", "French", "Indian", "Italian", "Japanese", "Mexican", "Thai", "Filipino", "Asian"}[c]
}

type Currencies int

const (
	Currency0 Currencies = iota
	Currency1
	Currency2
	Currency3
	Currency4
	Currency5
	Currency6
	Currency7
	Currency8
)

func (c Currencies) String() string {
	if c < Currency0 || c > Currency8 {
		return "USD"
	}
	return [...]string{"USD", "EUR", "GBP", "JPY", "CNY", "INR", "MXN", "THB", "PHP"}[c]
}

type TimeUnits int

const (
	TimeUnit0 TimeUnits = iota
	TimeUnit1
	TimeUnit2
	TimeUnit3
	TimeUnit4
)

func (t TimeUnits) String() string {
	if t < TimeUnit0 || t > TimeUnit4 {
		return "Unknown"
	}
	return [...]string{"Seconds", "Minutes", "Hours", "Days", "Weeks"}[t]
}

type RecipeTags int

const (
	RecipeTag0 RecipeTags = iota
	RecipeTag1
	RecipeTag2
	RecipeTag3
	RecipeTag4
	RecipeTag5
	RecipeTag6
	RecipeTag7
	RecipeTag8
	RecipeTag9
	RecipeTag10
	RecipeTag11
	RecipeTag12
	RecipeTag13
	RecipeTag14
	RecipeTag15
)

func (r RecipeTags) String() string {
	if r < RecipeTag0 || r > RecipeTag15 {
		return "Unknown"
	}
	return [...]string{"Main Course", "Side Dish", "Dessert", "Appetizer", "Salad", "Bread", "Breakfast", "Soup", "Beverage", "Sauce", "Marinade", "Preserves", "Spice Mix", "Dinner", "Lunch", "Snack"}[r]
}

type ShoppingTypes int

const (
	ShoppingType0 ShoppingTypes = iota
	ShoppingType1
	ShoppingType2
	ShoppingType3
)

func (s ShoppingTypes) String() string {
	if s < ShoppingType0 || s > ShoppingType3 {
		return "Unknown"
	}
	return [...]string{"Online", "In-person", "Delivery", "Pickup"}[s]
}
