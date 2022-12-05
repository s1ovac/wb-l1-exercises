package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func main() {
	stopByContext(context.Background())
	stopByChan()
	stopByWaitGroup()
}

func stopByContext(ctx context.Context) {
	ctx, _ = context.WithTimeout(ctx, time.Second*3)
	check := make(chan bool)
	go func() {
		for {
			select {
			case <-time.Tick(time.Millisecond * 500): // Посылает сигнал в канал каждые 500 мс и сообщает, что горутина жива
				fmt.Println("Goroutine is alive!")
			case <-ctx.Done(): // ожидаем получения данных в канал по истечению срока 3 с
				check <- true
				fmt.Println("Goroutine is dead while context is stopped!")
				return
			}
		}
	}()
	<-check

}

func stopByChan() {
	done := make(chan bool)
	tm := time.After(time.Second * 3) // Инициализирую таймер, который пошлет сигнал при завершении в канал tm
	go func() {
		for {
			select {
			case <-time.Tick(time.Millisecond * 500):
				fmt.Println("Goroutine is alive!")
			case <-tm:
				fmt.Println("Closing goroutine")
				done <- true // Посылаем данные в канал done
				close(done)  // Закрываем канал done
			case <-done:
				fmt.Println("Goroutine is dead while channel is closed")
				return
			}
		}
	}()
	<-done
}

func stopByWaitGroup() {
	wg := &sync.WaitGroup{}
	check := make(chan bool)
	wg.Add(1)
	go func() {
		for {
			foo, ok := <-check // Тут реализация с получением данных из канала, когда получим данные, тогда выйдем из фукнции
			if !ok {
				fmt.Println("Goroutine is dead by WaitGroup")
				wg.Done()
				return
			}
			fmt.Println(foo)
		}
	}()
	check <- true
	close(check)
	wg.Wait() // Ждем завершения всех горутин
}
