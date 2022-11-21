package main

import (
	"fmt"
	"math"
	"sort"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
func main() {
	tempratures := make(map[int][]float64)                                // Результирующая мапа
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Инициализируем значения для примера
	sort.Float64s(temps)                                                  // Предварительная сортировка
	for i := 0; i < len(temps); i++ {
		key := findKey(temps[i])                              // Ищем ключ для значений
		tempratures[key] = append(tempratures[key], temps[i]) // Добавляем температурные значения в слайс
	}
	fmt.Println(tempratures)
}

func findKey(n float64) int {
	tmp := int(math.Round(n))
	return tmp - (tmp % 10)
}
