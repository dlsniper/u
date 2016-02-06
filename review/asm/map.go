package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := make(map[int]int, 10)
	fmt.Printf("%#v", reflect.ValueOf(a).Pointer())
}