package Shapes

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	areaFormula := s.Side * s.Side
	return areaFormula
}

func (s Square) Perimeter() float64 {
	areaFormula := 4 * s.Side
	return areaFormula
}
