package main

import "sync"

// Реализовать конкурентную запись данных в map.

type Counter struct {
	mu *sync.Mutex // Используем мьютекс для синхронизации и для предоствращения гонок
	m  map[int]struct{}
	mp *sync.Map
}

// NewCounter - конструктор счетчика
func NewCounter() *Counter {
	return &Counter{
		m:  make(map[int]struct{}),
		mu: new(sync.Mutex),
		mp: new(sync.Map),
	}
}

func main() {
	cnt := NewCounter()
	for i := 0; i < 1000; i++ {
		go func(i int) {
			cnt.SendValueByMutex(i)
		}(i)
	}
	for i := 1000; i < 2000; i++ {
		go func(i int) {
			cnt.SendValueByMap(i)
		}(i)
	}

}

func (c *Counter) SendValueByMutex(value int) {
	c.mu.Lock()
	c.m[value] = struct{}{} // Пустая структура ничего не весит
	c.mu.Unlock()           // без defer, так как будет небольшой оверхед
}

// Обертка стандартной библиотеки на Mutex
func (c *Counter) SendValueByMap(value int) {
	c.mp.Store(value, struct{}{})
}
