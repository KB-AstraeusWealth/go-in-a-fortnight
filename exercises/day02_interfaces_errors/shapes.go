package day2

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ R float64 }

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) Perimeter() float64 {
	return math.Pi * (c.R + c.R)
}

type Rectangle struct{ W, H float64 }

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (r Rectangle) Perimeter() float64 {
	return (r.W + r.H) * 2
}

// Describe returns a human string for a Shape.
// HINT: type switch — `switch v := s.(type) { case Circle: ...; case Rectangle: ...; default: }`.
// It's NOT exhaustiveness-checked, so keep the default. (Add the "fmt" import.)
func Describe(s Shape) string {
	switch s.(type) {
	case Circle:
		return fmt.Sprintf("circle r=%v", s.(Circle).R)
	case Rectangle:
		return fmt.Sprintf("rectangle %vx%v", s.(Rectangle).W, s.(Rectangle).H)
	default:
		return fmt.Sprintf("unknown type: %T", s)
	}
}
