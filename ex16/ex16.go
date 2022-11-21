package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	quickSort(arr)
}

func quickSort(inputArr []int) {
	fmt.Println("Неотсортированный массив встроенным методом языка: ", inputArr)
	sort.Ints(inputArr) // Сортируем слайс интов встроенным пакетом sort, также имеются методы Float64s, Strings
	fmt.Println("Отсортированный массив встроенным методом языка: ", inputArr)

	fmt.Println("Сортируем свой тип")

	personArr := []Person{{"Ilya", 21}, {"Ann", 17}, {"John", 43}, {"Alex", 35}}
	fmt.Println("Неотсортированный массив собственного типа встроенным методом языка: ", personArr)
	sort.Slice(personArr, func(i, j int) bool {
		return personArr[i].Name < personArr[j].Name
	})
	fmt.Println("Отсортированный массив собственного типа встроенным методом языка: ", personArr)
}
