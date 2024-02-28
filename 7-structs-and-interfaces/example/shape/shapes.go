package Shapes

import "math"

type Circle struct {
	Radius float64
	Name   string
}

func NewCircle(radius float64) Circle {
	return Circle{Radius: radius, Name: "Circle"}
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) GetName() string {
	return c.Name
}

type Square struct {
	Side float64
	Name string
}

func NewSquare(side float64) Square {
	return Square{Side: side, Name: "Square"}
}
func (s Square) Area() float64 {
	return s.Side * s.Side
}
func (s Square) GetName() string {
	return s.Name
}

type Rectangle struct {
	Width  float64
	Height float64
	Name   string
}

func NewRectangle(width, height float64) Rectangle {
	return Rectangle{Width: width, Height: height, Name: "Rectangle"}
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) GetName() string {
	return r.Name
}

type Shape interface {
	GetName() string
	Area() float64
}

func GetTotalAreaOfShapes(shapes ...Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
