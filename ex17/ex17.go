package main

import "fmt"

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
		low := 0
		high := len(arr) - 1
		for low <= high {
			mid := (low + high) / 2
			guess := arr[mid]
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
