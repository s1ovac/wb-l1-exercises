package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {
	sleepByTime(time.Second * 5)
	sleepByContext(time.Second * 5)
}

// Реализация sleep через фукнцию After пакета time, которая возвращает канал, тоже самое, что и NewTimer
func sleepByTime(tm time.Duration) {
	fmt.Printf("Sleeping by Timer for %.2f sec.", tm.Seconds())
	until := time.After(tm)
	select {
	case <-until:
		return
	}
}

// Реализация sleep через контекст, который ждет завершения и разблокирует горутину
func sleepByContext(tm time.Duration) {
	fmt.Printf("Sleeping by Context for %.2f sec.", tm.Seconds())
	ctx, _ := context.WithTimeout(context.Background(), tm)
	<-ctx.Done()

}
