# How to define enums

This package contains enums used throughout the butler-ai backend. Enumerations are a set of named values that represent a specific type of data (think a distinct set of elements), making the code more readable and maintainable.


Define each enum as follows:

#### `enum/{{functional_area}}.go`
```go
package enum

type Example int

const (
    option1 Example = iota
    option2
    option3
)
```

#### `enum/{{functional_area}}_helper.go`
```go
package enum


func (e enum.Example) String() string {
    if e < enum.option1 || e > enum.option3 {
        return "Unknown"
    }
    return [...]string{"Option 1", "Option 2", "Option 3"}[e]
}
```

This way, the `Example` type now behaves as an enum, can be called in the codebase using the identifier `enum.Example`

**When defining/editing enums ensure function is mapping the enums to the right values**

