package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	cnt1, cnt2 := NewCounter(), NewCounter()
	for i := 0; i < 1000; i++ {
		go cnt1.IncViaMutex()
	}

	for i := 0; i < 1000; i++ {
		go cnt2.IncViaAtomic()
	}
	//runtime.Gosched() //Завершаем оставшиеся горутины
	fmt.Println("Результат счетчика через mutex", cnt1.Value())
	fmt.Println("Результат счетчика через atomic", cnt2.Value())
}

type Counter struct {
	mu *sync.Mutex
	c  int64
}

func NewCounter() *Counter {
	return &Counter{
		mu: &sync.Mutex{},
	}
}

func (c *Counter) IncViaMutex() {
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}

func (c *Counter) Value() int64 {
	return c.c
}

func (c *Counter) IncViaAtomic() {
	atomic.AddInt64(&c.c, 1)
}
