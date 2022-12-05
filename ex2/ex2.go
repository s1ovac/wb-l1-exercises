package main

import (
	"fmt"
	"sync"
	"time"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

func main() {
	array := [...]int{2, 4, 6, 8, 10}
	fmt.Println("--------------------BY TIME.SLEEP--------------------")
	for _, n := range array {
		go powPrint(n)
	}
	time.Sleep(time.Second)

	wg := &sync.WaitGroup{}
	fmt.Println("--------------------SYNC BY SYNC.WAITGROUP--------------------")
	for _, n := range array {
		wg.Add(1)
		go func(n int) {
			powPrint(n)
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println("--------------------SYNC BY CHANNEL--------------------")
	//Создаем буффериованный канал с размером массива array
	ch := make(chan int, len(array))

	for i, n := range array {
		go func(i, n int, ch chan<- int) {
			powPrint(n)
			ch <- i
		}(i, n, ch)
	}
	//Читаем из канала все данные, для завершения работ всех запущенных горутин
	for i := 0; i < len(array); i++ {
		<-ch
	}

}

func powPrint(n int) {
	fmt.Printf("Квадрат числа: %d = %d\n", n, n*n)
}
