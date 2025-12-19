package Shapes

import "fmt"

func Example() {
	shapes := []Shape{
		Rectangle{
			Width:  5.2,
			Height: 3,
		},
		Circle{
			Radius: 4,
		},
		Square{
			Side: 6.8,
		},
		Triangle{
			A: 3,
			B: 4,
			C: 5,
		},
	}

	for i := 0; i < len(shapes); i++ {
		fmt.Println("Area", shapes[i].Area())
		fmt.Println("Perimeter", shapes[i].Perimeter())
	}
}
