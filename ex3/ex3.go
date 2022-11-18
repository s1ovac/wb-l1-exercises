package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

func main() {
	array := [...]int{2, 4, 6, 8, 10}
	var sum int
	fmt.Println("--------------------BY CHANNEL--------------------")
	//Создаем буффериованный канал с размером массива array
	ch := make(chan int, len(array))

	for _, n := range array {
		go func(n int, ch chan<- int) {
			ch <- n * n
		}(n, ch)
	}
	//Читаем из канала все данные, для завершения работ всех запущенных горутин
	for i := 0; i < len(array); i++ {
		sum += <-ch
	}
	fmt.Printf("Сумма квадратов чисел: %d\n", sum)
	sum = 0
	fmt.Println("--------------------BY SYNC.WAITGROUP--------------------")
	wg := &sync.WaitGroup{}
	for _, n := range array {
		wg.Add(1)
		go func(n int) {
			sum += n * n
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Printf("Сумма квадратов чисел: %d\n", sum)
}
