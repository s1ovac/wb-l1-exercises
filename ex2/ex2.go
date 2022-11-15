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

	for _, n := range array {
		go pow(n)
	}
	time.Sleep(time.Second)
	fmt.Println("--------------------BY TIME.SLEEP--------------------")
	wg := sync.WaitGroup{}
	for _, n := range array {
		wg.Add(1)
		go func(n int) {
			pow(n)
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println("--------------------SYNC BY SYNC.WAITGROUP--------------------")

}

func pow(n int) {
	fmt.Println(n * n)
}
