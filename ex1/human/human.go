package human

import "fmt"

type Human struct {
	//Имеем поля структуры Human
	name string
	age  int
	//...
}

func (h *Human) SayHello(times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Hello")
	}
}
