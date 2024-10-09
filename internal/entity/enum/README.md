# How to define enums

This package contains enums used throughout the butler-ai backend. Enumerations are a set of named values that represent a specific type of data (think a distinct set of elements), making the code more readable and maintainable.

Define each enum as follows:

#### `enum/enums_n.go`
```go
package enum

type Examples int

const (
    Example0 Examples = iota
    Example1
    Example2
)

func (e enum.Example) String() string {
    if e < enum.option1 || e > enum.option3 {
        return "Unknown"
    }
    return [...]string{"Option 1", "Option 2", "Option 3"}[e]
}
```

This way, the `Examples` type now behaves as an enum, can be called in the codebase using the identifier `enum.Examples`

### List of available enums:

#### [`enums_1.go`](enums_1.go) :
- Genders,
- HeightUnits,
- WeightUnits,
- Lifestyles,
- ChatTypes,
- Cuisines,
- Currencies,
- TimeUnits,
- RecipeTags,
- ShoppingTypes

#### [`enums_2.go`](enums_2.go) :
- Difficulties,
- GoalTypes,
- MealTypes,
- NutritionTags
