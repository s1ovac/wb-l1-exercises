package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
//
//	aabcd — false
func main() {
	testString := "abcdEfA"
	if Distinct(testString) {
		fmt.Printf("String '%s' has only distinct symbols\n", testString)
		return
	}
	fmt.Printf("String '%s' has not distinct symbols\n", testString)
}

func Distinct(str string) bool {
	str = strings.ToLower(str)
	distMap := make(map[string]int)
	for _, s := range str {
		distMap[string(s)] += 1
		if distMap[string(s)] > 1 {
			return false
		}
	}
	return true
}
