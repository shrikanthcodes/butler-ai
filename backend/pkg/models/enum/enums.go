package enum

type Gender int

const (
	Unknown Gender = iota
	Male
	Female
	NonBinary
)

type HeightUnit int

const (
	Centimeters HeightUnit = iota
	Inches
)

type WeightUnit int

const (
	Kilograms WeightUnit = iota
	Pounds
)

type Lifestyle int

const (
	Sedentary Lifestyle = iota
	Average
	Active
)

type ChatType int

const (
	RecipeChat ChatType = iota
	ShoppingChat
	HealthChat
	MotivationChat
	CalorieTrackerChat
)

type Cuisines int

const (
	African Cuisines = iota
	American
	Chinese
	French
	Indian
	Italian
	Japanese
	Mexican
	Thai
	Filipino
	Asian
)

type Currencies int

const (
	USD Currencies = iota
	EUR
	GBP
	JPY
	CNY
	INR
	MXN
	THB
	PHP
)

type TimeUnits int

const (
	Seconds TimeUnits = iota
	Minutes
	Hours
	Days
	Weeks
)

type RecipeTags int

const (
	MainCourse RecipeTags = iota
	SideDish
	Dessert
	Appetizer
	Salad
	Bread
	Breakfast
	Soup
	Beverage
	Sauce
	Marinade
	Preserves
	SpiceMix
	Dinner
	Lunch
	Snack
)

type ShopMode int

const (
	Online ShopMode = iota
	InPerson
	Delivery
	Pickup
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type GoalTypes int

const (
	WeightLoss GoalTypes = iota
	WeightGain
	SelfLove
	Addiction
	Custom
	NewSkill
	Discipline
)
