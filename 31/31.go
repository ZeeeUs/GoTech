package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func calcDistance(p1, p2 Point) float64 {
	res := math.Sqrt(math.Pow(p2.X - p1.X, 2)+math.Pow(p2.Y-p1.Y, 2))
	return res
}

func main() {
	p1 := Point{X: 2, Y: 4}
	p2 := Point{X: 3, Y: 5}

	fmt.Println(calcDistance(p1, p2))
}
