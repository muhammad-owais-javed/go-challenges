package sprint

type Circle struct {

	Radius float32
	Diameter float32
	Area float32
	Perimeter float32

}

func GetCircle(r float32) Circle {

 var c Circle
 var p float32 = 3.14
 c.Radius = r
 c.Diameter = r * 2
 c.Area = (p*(r * r))
 c.Perimeter = c.Diameter * p

return c

}