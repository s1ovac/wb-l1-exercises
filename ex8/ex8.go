package main

import "fmt"

func main() {
	var number int64
	var i int64
	fmt.Print("Введите число: ") // Бит   16 8 4 2 1 0
	fmt.Scan(&number)            // Число X  X X X X
	fmt.Print("Введите индекс: ")
	fmt.Scan(&i)
	fmt.Printf("Переменная в int64: %d\nПеременная в двоичной форме: %b\n", number, number)
	temp := 1 << i
	fmt.Printf("Полученная переменная в int64: %d\nПолученная переменная в двоичной форме: %b\n", number^int64(temp), number^int64(temp))
}
