package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func distance(p1 Point, p2 Point) int {
	return int(math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))))
}

func main() {
	a := NewPoint(0, 1)
	b := NewPoint(2, 7)

	dist := distance(a, b)

	fmt.Println("distance from A to B: ", dist)
}
