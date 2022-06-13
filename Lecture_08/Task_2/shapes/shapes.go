package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Shapes []Shape

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func NewRectangle(height, width float64) Rectangle {
	return Rectangle{Width: width, Height: height }
}

func NewCircle(radius float64) Circle {
	return Circle{Radius: radius}
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (circ Circle) Area() float64 {
	return math.Pi * circ.Radius * circ.Radius
}

func PrintArea(shape Shape) {
	fmt.Println(shape.Area())
}

func (shapes Shapes) LargestArea() float64{
	maxValue := shapes[0].Area()
	for idx, _ := range shapes {
		if maxValue < shapes[idx].Area(){
			maxValue = shapes[idx].Area()
		}
	}
	return maxValue
}