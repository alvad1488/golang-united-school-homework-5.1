package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (inp Square) End() Point {
	var (
		sPtr Point
		side int
		ePtr Point
	)
	sPtr = inp.start
	side = int(inp.a)
	ePtr.x = sPtr.x + side
	ePtr.y = sPtr.y + side

	return ePtr

}

func (input Square) Area() uint {
	var (
		side uint
		area uint
	)

	side = input.a
	area = uint(math.Pow(float64(side), 2))

	return area
}

func (input Square) Perimeter() uint {
	var (
		side uint
		per  uint
	)
	side = input.a
	per = side * 4

	return per
}
