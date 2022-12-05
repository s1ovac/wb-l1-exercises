package main

import (
	"fmt"
	"reflect"
)

// Разработать программу,
// которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
func main() {
	// Инициализируем переменную, которую хоти проверить
	var a = make(chan int)
	checkType(a)
	checkTypeByReflect(a)
	checkTypeByFmt(a)
}

func checkType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("Type of %v is 'int'\n", a)
	case string:
		fmt.Printf("Type of %v is 'string'\n", a)
	case chan bool:
		fmt.Printf("Type of %v is 'boolean'\n", a)
	case chan int:
		fmt.Printf("Type of %v is 'chan int'\n", a)
	case chan float64:
		fmt.Printf("Type of %v is 'chan float'\n", a)
	case chan string:
		fmt.Printf("Type of %v is 'chan string'\n", a)
	case bool:
		fmt.Printf("Type of %v is 'chan bool'\n", a)
	default:
		fmt.Printf("Unknown type")
	}
}

func checkTypeByReflect(a interface{}) {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		fmt.Printf("Type of %v is 'int'\n", a)
	case reflect.String:
		fmt.Printf("Type of %v is 'string'\n", a)
	case reflect.Bool:
		fmt.Printf("Type of %v is 'boolean'\n", a)
	case reflect.Chan:
		fmt.Printf("Type of %v is 'chan'\n", a)
	default:
		fmt.Printf("Unknown type")
	}
}

func checkTypeByFmt(a interface{}) {
	fmt.Printf("Type of %v is  %T\n", a, a)

}
