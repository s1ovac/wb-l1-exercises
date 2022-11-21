package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	subset := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Собственное множество: ", newSub(subset))

}

func newSub(arr []string) []string {
	subM := make(map[string]struct{}) //Инициализация легковесной мапы
	subArr := make([]string, 0)       //Инициализация слайса

	for _, val := range arr { //Добавляем ключи в мапу из слайса
		subM[val] = struct{}{}

	}

	for key := range subM { //Создаем новое множество
		subArr = append(subArr, key)
	}

	return subArr
}
