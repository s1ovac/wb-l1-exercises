package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров

func main() {
	var (
		workers int
		wg      = new(sync.WaitGroup)
		flag    bool
	)
	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan int)
	ticker := time.NewTicker(time.Second)
	fmt.Print("Введите кол-во воркеров: ")
	if _, err := fmt.Scan(&workers); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < workers; i++ {
		go work(ctx, wg, input, i)
	}
	go shutdown(cancel)
	wg.Wait()

	for {
		select {
		case <-ticker.C:
			input <- time.Now().Second()
		case <-ctx.Done():
			flag = true
			break
		}
		if flag {
			break
		}

	}
}

func work(ctx context.Context, wg *sync.WaitGroup, input chan int, i int) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine with ID: %d is Done\n", i) // для проверки, что все горутины завершились
			return
		case x := <-input:
			fmt.Printf("Goroutine with ID: %d Data is: %v\n", i, x)
		}
	}
}

func shutdown(cancel context.CancelFunc) {
	chCtx := make(chan os.Signal, 1)       //создаём канал для сигналов системы
	signal.Notify(chCtx, os.Interrupt)     //ждем сигнал ОС что наше приложение завершилось
	log.Printf("system call:%+v", <-chCtx) // выводим полученный сигнал в stdout

	cancel() // отменяем контекст
}
