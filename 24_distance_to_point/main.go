package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p1 *Point) DistanceToPoint(p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	p1 := &Point{}
	fmt.Println("Введите координаты X и Y первой точки")
	fmt.Scan(&p1.x, &p1.y)

	p2 := &Point{}
	fmt.Println("Введите координаты X и Y второй точки")
	fmt.Scan(&p2.x, &p2.y)

	fmt.Printf("Согласно моим вычислениям, расстояние между ними %f", p1.DistanceToPoint(p2))
}
