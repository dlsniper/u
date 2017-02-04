package main

import (
	"fmt"
)

func main() {
	a := "World"
	b := int(123)
	c := struct {
		Demo  bool
		Demo2 float32
	}{Demo: true, Demo2: 3.14}
	d := "World"
	e := "World"
	fmt.Printf("Hello %s %s %s %s %s\n", a, b, c, d, e)
}
