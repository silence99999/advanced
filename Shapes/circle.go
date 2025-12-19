package Shapes

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	areaFormula := math.Pi * (c.Radius * c.Radius)
	return areaFormula
}

func (c Circle) Perimeter() float64 {
	perimeterFormula := 2 * math.Pi * c.Radius
	return perimeterFormula
}
