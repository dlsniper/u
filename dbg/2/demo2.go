package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	a := struct {
		A string
		B int
	}{A: "demo", B: 10}
	fmt.Printf("%#v\n", a)
	b := map[string]string{"demo": "demo"}
	fmt.Printf("%#v\n", b)
	c := []string{"demo"}
	fmt.Printf("%#v\n", c)
	d := []int{1}
	fmt.Printf("%#v\n", d)
	e := map[int]int{1: 1}
	fmt.Printf("%#v\n", e)
	var f uint64 = 11
	fmt.Printf("%#v\n", f)
	g := true
	fmt.Printf("%#v\n", g)

	z := make(chan int, 10)
	go fibonacci(cap(z), z)
	for i := range z {
		fmt.Println(i)
	}

	time.Sleep(30 * time.Second)
	fmt.Println("out of time")

	z1 := make(chan int, 10)
	fibonacci(cap(z1), z1)
	for i := range z1 {
		fmt.Println(i)
	}
}
