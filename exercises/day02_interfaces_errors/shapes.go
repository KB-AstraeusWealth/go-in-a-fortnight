package day2

// Implement the Area/Perimeter methods and Describe. Add imports as you go.

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ R float64 }

// HINT: math.Pi for the circle (add the "math" import).
func (c Circle) Area() float64      { panic("TODO: implement Circle.Area") }
func (c Circle) Perimeter() float64 { panic("TODO: implement Circle.Perimeter") }

type Rectangle struct{ W, H float64 }

func (r Rectangle) Area() float64      { panic("TODO: implement Rectangle.Area") }
func (r Rectangle) Perimeter() float64 { panic("TODO: implement Rectangle.Perimeter") }

// Describe returns a human string for a Shape.
// HINT: type switch — `switch v := s.(type) { case Circle: ...; case Rectangle: ...; default: }`.
// It's NOT exhaustiveness-checked, so keep the default. (Add the "fmt" import.)
func Describe(s Shape) string { panic("TODO: implement Describe with a type switch") }
