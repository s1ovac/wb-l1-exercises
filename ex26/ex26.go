package main

import "fmt"

func main() {
	test := "abcdef"
	fmt.Println(Distinct(test))
}

func Distinct(str string) bool {
	distMap := make(map[string]int)
	for _, s := range str {
		distMap[string(s)] += 1
		if distMap[string(s)] > 1 {
			return false
		}
	}
	return true
}
