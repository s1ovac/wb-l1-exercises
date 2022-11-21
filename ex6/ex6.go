package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.Gosched()
	goroutineStopByChannel()
	goroutineStopByContext()
}

func goroutineStopByChannel() {
	done := make(chan bool)
	// Чтобы сделать goroutine остановливаемой, позволим ей слушать сигнал остановки на канале.
	go func() { //запускаю горутины
		for { //в цикле слушаю два канала
			select {
			case <-time.Tick(time.Millisecond * 500): //каждую секунду вывожу сообщение
				fmt.Println("Горутина активна в канале")
			case <-done: //жду сигнал о завершении от другой горутины
				return
			}
		}
	}()
	go func() {
		time.Sleep(time.Second * 2)
		done <- true
		close(done)
	}()

	fmt.Println("Активных горутин: ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second) //даю горутинам время на отработку

	fmt.Println("Активных горутин: ", runtime.NumGoroutine())
	fmt.Println("Горутина остановлена")
	fmt.Println("Выход программы")
}

func goroutineStopByContext() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-time.Tick(500 * time.Millisecond):
				fmt.Println("Горутина активна в контексте")
			case <-ctx.Done():
				fmt.Println("Горутина остановалена в контексте")
				done <- true
				return
			}

		}
	}()
	fmt.Println("Активных горутин: ", runtime.NumGoroutine())
	<-done
	fmt.Println("Активных горутин: ", runtime.NumGoroutine())
	fmt.Println("Горутина остановлена")
	fmt.Println("Выход программы")
}
