package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

var justString string

func main() {
	someFunc()
	correctFunc()
}

func createHugeString(size int) (result string) {
	var str strings.Builder
	for i := 0; i < size; i++ {
		str.WriteString("привет ")
	}
	return str.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
}

func correctFunc() {
	v := createHugeString(1 << 10)
	tmp := []rune(v)
	justString = string(tmp[:100])
	fmt.Println(justString)
}

// Получим корректный ответ - привет привет привет привет привет привет привет привет привет привет привет привет привет привет пр

// Негативные последствия - это потеря данных, т.к unicode символы занимают 2 байта и получим потерю информации
// привет привет привет привет привет привет привет прив� - пример 50 символов, а должно быть 100
