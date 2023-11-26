package main

import (
	"fmt"
	"math"
)

// Point Структура Point представляет точку с координатами x и y.
type Point struct {
	x, y float64
}

// NewPoint Конструктор NewPoint создает новый экземпляр структуры Point.
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Distance Метод Distance вычисляет расстояние между текущей точкой и другой точкой.
func (p *Point) Distance(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки с помощью конструктора NewPoint.
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	// Вычисляем расстояние между точками.
	distance := point1.Distance(point2)

	// Выводим результат.
	fmt.Printf("Расстояние между точкой 1 и точкой 2: %.2f\n", distance)
}
