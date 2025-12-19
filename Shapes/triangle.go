package Shapes

import (
	"fmt"
	"math"
)

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	semiP := t.Perimeter() / 2
	heronFormula := semiP * (semiP - t.A) * (semiP - t.B) * (semiP - t.C)
	res := math.Sqrt(heronFormula)
	if res == math.NaN() {
		fmt.Println("Please enter valid sides")
		return -1
	}
	return res

}

func (t Triangle) Perimeter() float64 {
	perimeterFormula := t.A + t.B + t.C
	return perimeterFormula
}
