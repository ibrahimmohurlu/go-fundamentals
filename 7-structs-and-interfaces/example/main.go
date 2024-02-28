package main

import (
	"fmt"
	Shapes "structs-and-interfaces/example/shape"
)

func main() {

	circle := Shapes.NewCircle(3.72)
	rectangle := Shapes.NewRectangle(7.54, 4.45)
	square := Shapes.NewSquare(6.66)
	shapes := []Shapes.Shape{circle, rectangle, square}

	for _, shape := range shapes {
		fmt.Printf("Area of %v is %f \n", shape.GetName(), shape.Area())
	}

	fmt.Printf("Total area of shapes:%f\n", Shapes.GetTotalAreaOfShapes(shapes...))
}
