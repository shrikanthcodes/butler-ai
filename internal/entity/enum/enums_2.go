package enum

type Difficulties int

const (
	Difficulties0 Difficulties = iota
	Difficulties1
	Difficulties2
)

func (d Difficulties) String() string {
	if d < Difficulties0 || d > Difficulties2 {
		return "Unknown"
	}
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

type GoalTypes int

const (
	GoalTypes0 GoalTypes = iota
	GoalTypes1
	GoalTypes2
	GoalTypes3
	GoalTypes4
	GoalTypes5
	GoalTypes6
)

func (g GoalTypes) String() string {
	if g < GoalTypes0 || g > GoalTypes6 {
		return "Unknown"
	}
	return [...]string{"Weight Loss", "Weight Gain", "Self Love", "Addiction", "Custom", "New Skill", "Discipline"}[g]
}

type MealTypes int

const (
	MealTypes0 MealTypes = iota
	MealTypes1
	MealTypes2
	MealTypes3
)

func (m MealTypes) String() string {
	if m < MealTypes0 || m > MealTypes3 {
		return "Unknown"
	}
	return [...]string{"Breakfast", "Lunch", "Dinner", "Snack"}[m]
}

type NutritionTags int

const (
	NutritionTags0 NutritionTags = iota
	NutritionTags1
	NutritionTags2
	NutritionTags3
	NutritionTags4
	NutritionTags5
	NutritionTags6
	NutritionTags7
)

func (n NutritionTags) String() string {
	if n < NutritionTags0 || n > NutritionTags7 {
		return "Unknown"
	}
	return [...]string{"Low Calorie", "High Protein", "Low Carb", "Low Fat", "High Fiber", "Low Sodium", "Low Sugar", "Balanced"}[n]
}

type Roles int

const (
	Roles0 Roles = iota
	Roles1
)

func (r Roles) String() string {
	if r < Roles0 || r > Roles1 {
		return "Unknown"
	}
	return [...]string{"User", "Admin"}[r]
}
