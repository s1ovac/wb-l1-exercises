package main

import (
	"fmt"
	"strings"
)

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
