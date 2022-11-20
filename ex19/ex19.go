package main

import (
	"fmt"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func main() {
	var check string
	fmt.Scan(&check)
	fmt.Printf("Original string: %s\nReverse string: %s\n", check, ReverseString(check))
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
