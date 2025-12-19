package Shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	areaFormula := r.Width * r.Height
	return areaFormula
}

func (r Rectangle) Perimeter() float64 {
	perimeterFormula := 2 * (r.Width + r.Height)
	return perimeterFormula
}
