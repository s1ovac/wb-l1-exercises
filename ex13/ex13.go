package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)
}
