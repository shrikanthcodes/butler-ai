package enum

type Genders int

const (
	Genders0 Genders = iota
	Genders1
	Genders2
	Genders3
)

func (g Genders) String() string {
	if g < Genders0 || g > Genders3 {
		return "Unknown"
	}
	return [...]string{"Unknown", "Male", "Female", "Non-Binary"}[g]
}

type HeightUnits int

const (
	HeightUnits0 HeightUnits = iota
	HeightUnits1
)

func (h HeightUnits) String() string {
	if h < HeightUnits0 || h > HeightUnits1 {
		return "Inches"
	}
	return [...]string{"Centimeters", "Inches"}[h]
}

type WeightUnits int

const (
	WeightUnits0 WeightUnits = iota
	WeightUnits1
)

func (w WeightUnits) String() string {
	if w < WeightUnits0 || w > WeightUnits1 {
		return "Pounds"
	}
	return [...]string{"Kilograms", "Pounds"}[w]
}

type Lifestyles int

const (
	Lifestyles0 Lifestyles = iota
	Lifestyles1
	Lifestyles2
)

func (l Lifestyles) String() string {
	if l < Lifestyles0 || l > Lifestyles2 {
		return "Medium"
	}
	return [...]string{"Sedentary", "Average", "Active"}[l]
}

type ChatTypes int

const (
	ChatTypes0 ChatTypes = iota
	ChatTypes1
	ChatTypes2
	ChatTypes3
	ChatTypes4
)

func (c ChatTypes) String() string {
	if c < ChatTypes0 || c > ChatTypes4 {
		return "Unknown"
	}
	return [...]string{"Recipe", "Shopping", "Health", "Motivation", "Calorie Tracker"}[c]
}

type Cuisines int

const (
	Cuisines0 Cuisines = iota
	Cuisines1
	Cuisines2
	Cuisines3
	Cuisines4
	Cuisines5
	Cuisines6
	Cuisines7
	Cuisines8
	Cuisines9
	Cuisines10
)

func (c Cuisines) String() string {
	if c < Cuisines0 || c > Cuisines10 {
		return "Unknown"
	}
	return [...]string{"African", "American", "Chinese", "French", "Indian", "Italian", "Japanese", "Mexican", "Thai", "Filipino", "Asian"}[c]
}

type Currencies int

const (
	Currencies0 Currencies = iota
	Currencies1
	Currencies2
	Currencies3
	Currencies4
	Currencies5
	Currencies6
	Currencies7
	Currencies8
)

func (c Currencies) String() string {
	if c < Currencies0 || c > Currencies8 {
		return "USD"
	}
	return [...]string{"USD", "EUR", "GBP", "JPY", "CNY", "INR", "MXN", "THB", "PHP"}[c]
}

type TimeUnits int

const (
	TimeUnits0 TimeUnits = iota
	TimeUnits1
	TimeUnits2
	TimeUnits3
	TimeUnits4
)

func (t TimeUnits) String() string {
	if t < TimeUnits0 || t > TimeUnits4 {
		return "Unknown"
	}
	return [...]string{"Seconds", "Minutes", "Hours", "Days", "Weeks"}[t]
}

type RecipeTags int

const (
	RecipeTags0 RecipeTags = iota
	RecipeTags1
	RecipeTags2
	RecipeTags3
	RecipeTags4
	RecipeTags5
	RecipeTags6
	RecipeTags7
	RecipeTags8
	RecipeTags9
	RecipeTags10
	RecipeTags11
	RecipeTags12
	RecipeTags13
	RecipeTags14
	RecipeTags15
)

func (r RecipeTags) String() string {
	if r < RecipeTags0 || r > RecipeTags15 {
		return "Unknown"
	}
	return [...]string{"Main Course", "Side Dish", "Dessert", "Appetizer", "Salad", "Bread", "Breakfast", "Soup", "Beverage", "Sauce", "Marinade", "Preserves", "Spice Mix", "Dinner", "Lunch", "Snack"}[r]
}

type ShoppingTypes int

const (
	ShoppingTypes0 ShoppingTypes = iota
	ShoppingTypes1
	ShoppingTypes2
	ShoppingTypes3
)

func (s ShoppingTypes) String() string {
	if s < ShoppingTypes0 || s > ShoppingTypes3 {
		return "Unknown"
	}
	return [...]string{"Online", "In-person", "Delivery", "Pickup"}[s]
}
