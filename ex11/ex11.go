package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств
func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4, 5}
	Combine(a, b)
}

func Combine(a []int, b []int) {
	mp := make(map[int]int)
	resultArr := make([]int, 0)
	// Добавляем уникальные значения множества А в ключи мапы для дальнейшего поиска
	for _, v := range a {
		if _, ok := mp[v]; !ok {
			mp[v] += 1
		}
	}
	// Инкрементируем повторяющееся ключи, для составления множества
	for _, v := range b {
		mp[v] += 1
	}
	// Ищем по ключам мапы повторяющееся значения и состовляем результирующий массив
	for key, value := range mp {
		if value > 1 {
			resultArr = append(resultArr, key)
		}
	}
	fmt.Println(resultArr)
}
