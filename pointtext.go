package sprint

import "fmt"

type Point struct{

	X float32
	Y float32
	Text string
}

func PointText(p Point) Point {

	return Point{p.X, p.Y, fmt.Println("Text at (%f, %f):", p.X, p.Y)}
}