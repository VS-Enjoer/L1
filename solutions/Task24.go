package Solution

import (
	"fmt"
	"math"
)

/*
Задание 24:
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде
структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

// Point представляет точку с координатами x и y.

type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance вычисляет расстояние между двумя точками.
func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func Task24() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	distance := Distance(point1, point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
