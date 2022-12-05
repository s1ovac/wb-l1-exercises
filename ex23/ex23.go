package main

import "fmt"

// Удалить i-ый элемент из слайса.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(RemoveItem(2, slice))
	fmt.Println(RemoveItemByCopy(2, slice))
}

func RemoveItem(index int, slice []int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveItemByCopy(index int, slice []int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
