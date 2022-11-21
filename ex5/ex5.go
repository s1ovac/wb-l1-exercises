package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool) // Дополнительный канал с типом данных bool
	until := time.After(time.Second * 10)
	go send(msg, done)
	for {
		select {
		case m := <-msg:
			// Читаем данные из канала
			fmt.Println(m)
		case <-until:
			done <- true
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done: // Ждем получения данных для закрытия канала
			fmt.Println("done")
			close(ch)
			return
		default:
			// Пишем данные в канал
			ch <- "hello"
		}

	}
}
