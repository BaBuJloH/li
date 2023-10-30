// 24. Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"Level1/point"
	"math"
)

// "конструктор" возвращает указатель на структуру
func NewPoint(x, y float64) *point.Point {
	p := &point.Point{}
	p.Set(x, y)
	return p
}

func main() {
	a := NewPoint(4, 3)
	b := NewPoint(4, 8)

	dist := math.Sqrt(math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2))
	fmt.Println(dist)
}
