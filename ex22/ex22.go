package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b = new(big.Int), new(big.Int)
	var mode int8
	fmt.Println("Введите два числа:")
	fmt.Scan(a, b)
	fmt.Printf("Выберите операцию:\n1 - Сложениe\t2 - Вычитание\n3 - Умножение\t4 - Деление\n")
	fmt.Scan(&mode)
	switch mode {
	case 1:
		fmt.Printf("Результат сложения: %v\n", a.Add(a, b))
	case 2:
		fmt.Printf("Результат вычитания: %v\n", a.Sub(a, b))
	case 3:
		fmt.Printf("Результат умножения: %v\n", a.Mul(a, b))
	case 4:
		fmt.Printf("Результат деления: %v\n", a.Div(a, b))
	default:
		fmt.Println("Вы выбрали некорректный режим!")

	}
}
