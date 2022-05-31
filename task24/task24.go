package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y float64
}

func main() {
	A := NewPoint(1, 0)
	B := NewPoint(5, 5)
	fmt.Println("Расстояние между точками :", GetDistance(A, B))
}

func NewPoint(x, y float64) *Point {
	newStr := &Point{
		x: x,
		y: y,
	}
	return newStr
}

//собственно сама формула AB = √(xb - xa)^2 + (yb - ya)^2
func GetDistance(f, s *Point) float64 {
	return math.Sqrt((s.x-f.x)*(s.x-f.x) + (s.y-f.y)*(s.y-f.y))
}
