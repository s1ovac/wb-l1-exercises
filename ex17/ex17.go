package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var searchNum int
	fmt.Print("Enter number to find: ")
	fmt.Scan(&searchNum)
	if result := BinarySearch(arr, searchNum); result > -1 {
		fmt.Printf("Index of number %d is %d\n", searchNum, BinarySearch(arr, searchNum))
	} else {
		fmt.Printf("Sorry, Number %d not found in array\n", searchNum)
	}
}

// На вход должен идти отсортироанный слайс, возвращаем найденный индекс
func BinarySearch(arr []int, item int) int {
	if len(arr) > 0 {
		low := 0             // Нижний диапазон индекса
		high := len(arr) - 1 // Верхний диапазон индекса
		for low <= high {
			mid := (low + high) / 2 // Формула вычисления среднего индекса
			guess := arr[mid]       // Если не находим число, то ищем дальше.
			if guess == item {
				return mid
			}
			if guess > item {
				high = mid - 1
			}
			if guess < item {
				low = mid + 1
			}
		}
	}
	return -1
}
