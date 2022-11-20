package main

import (
	"fmt"

	"github.com/s1ovac/wv-l1-exercises/ex24/point"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
func main() {
	p1, p2 := new(point.Point), new(point.Point)
	// Выставляем координаты первой точки
	p1.SetCoordinates(1, 1)
	// Выставляем координаты второй точки
	p2.SetCoordinates(2, 2)
	// Вывод результата
	fmt.Printf("Расстояние между точками: %.2f\n", point.CountDestination(*p1, *p2))
}
