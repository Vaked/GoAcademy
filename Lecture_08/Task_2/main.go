package main

import (
	"fmt"
	shapes "shapes/shapes"
)

func main() {
	var shapesList shapes.Shapes

	shapesList = append(shapesList, shapes.NewRectangle(4,4))
	shapesList = append(shapesList, shapes.NewRectangle(2,3))
	shapesList = append(shapesList, shapes.NewRectangle(6,7))
	shapesList = append(shapesList, shapes.NewRectangle(8,9))

	shapesList = append(shapesList, shapes.NewCircle(2))
	shapesList = append(shapesList, shapes.NewCircle(3))
	shapesList = append(shapesList, shapes.NewCircle(4))

	fmt.Println(shapesList.LargestArea())
}