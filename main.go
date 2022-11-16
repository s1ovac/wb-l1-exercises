package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	s := sort.IntSlice{}
	fmt.Println(reflect.TypeOf(s))
	fmt.Printf("%T\n", s)
	s = append(s, 5)
}
