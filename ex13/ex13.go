package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)
	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)

	//побитовое исключающее (ИЛИ)
	a = a ^ b // 100^10 = 110(6)
	b = a ^ b // 110^10 = 100(4)
	a = a ^ b // 110^100 = 010(2)
	fmt.Printf("Побитовое исключающее: a= %d, b= %d\n", a, b)
}
