package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	mu := &sync.Mutex{}
	arr := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(arr))
	chPow := make(chan int, len(arr))
	for _, v := range arr {
		mu.Lock()
		ch <- v
		mu.Unlock()
	}
	close(ch)
	for v := range ch {
		mu.Lock()
		chPow <- v * v
		mu.Unlock()
	}
	close(chPow)
	for pow := range chPow {
		fmt.Println(pow)
	}
}
