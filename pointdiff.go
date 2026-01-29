package sprint

type Point struct{

	X float32
	Y float32
	Text string
}

func PointDiff(p1, p2 Point) Point {

	if p1.X > p2.Y {
		return p1
	} else if p2.Y > p1.X {
		return p2
	} else p1.X == p2.Y {
		return p2
	}
}