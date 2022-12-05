package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	mu := &sync.Mutex{}
	arr := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(arr))    // Канал в который будут записываться числа из массива.
	chPow := make(chan int, len(arr)) // Канал, в который будут записываться квадраты чисел
	for _, v := range arr {
		mu.Lock()
		ch <- v
		mu.Unlock()
	}
	close(ch) // Закрываем канал, так как не будем в него уже ничего писать
	for v := range ch {
		mu.Lock()
		chPow <- v * v
		mu.Unlock()
	}
	close(chPow) // Закрываем канал, так как не будем в него уже ничего писать
	for pow := range chPow {
		fmt.Println(pow)
	}
}
