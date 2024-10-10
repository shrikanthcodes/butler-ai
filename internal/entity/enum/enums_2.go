package enum

type Difficulties int

const (
	Difficulty0 Difficulties = iota
	Difficulty1
	Difficulty2
)

func (d Difficulties) String() string {
	if d < Difficulty0 || d > Difficulty2 {
		return "Unknown"
	}
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

type GoalTypes int

const (
	GoalType0 GoalTypes = iota
	GoalType1
	GoalType2
	GoalType3
	GoalType4
	GoalType5
	GoalType6
)

func (g GoalTypes) String() string {
	if g < GoalType0 || g > GoalType6 {
		return "Unknown"
	}
	return [...]string{"Weight Loss", "Weight Gain", "Self Love", "Addiction", "Custom", "New Skill", "Discipline"}[g]
}

type MealTypes int

const (
	MealType0 MealTypes = iota
	MealType1
	MealType2
	MealType3
)

func (m MealTypes) String() string {
	if m < MealType0 || m > MealType3 {
		return "Unknown"
	}
	return [...]string{"Breakfast", "Lunch", "Dinner", "Snack"}[m]
}

type NutritionTags int

const (
	NutritionTag0 NutritionTags = iota
	NutritionTag1
	NutritionTag2
	NutritionTag3
	NutritionTag4
	NutritionTag5
	NutritionTag6
	NutritionTag7
)

func (n NutritionTags) String() string {
	if n < NutritionTag0 || n > NutritionTag7 {
		return "Unknown"
	}
	return [...]string{"Low Calorie", "High Protein", "Low Carb", "Low Fat", "High Fiber", "Low Sodium", "Low Sugar", "Balanced"}[n]
}

type Roles int

const (
	Role0 Roles = iota
	Role1
)

func (r Roles) String() string {
	if r < Role0 || r > Role1 {
		return "Unknown"
	}
	return [...]string{"User", "Admin"}[r]
}
