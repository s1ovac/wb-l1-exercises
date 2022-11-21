package main

import (
	"fmt"
	"strings"
)

func main() {
	var words string = "snow dog sun"
	fmt.Println(ReverseWordsInString(words))
}

func ReverseWordsInString(words string) (result string) {
	words = strings.TrimSpace(words)
	split := strings.Split(words, " ")
	for i := 0; i < len(split)-2; i++ {
		split[i], split[i+2] = split[i+2], split[i]
	}
	return strings.Join(split, " ")
}
