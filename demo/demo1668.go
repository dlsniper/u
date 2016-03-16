package main

import "fmt"

type demo struct {
	Field string
}

func getMore() (demo1, demo2, demo3 *demo) {
	return &demo{"demo1"}, &demo{"demo2"}, &demo{"demo3"}
}

func main() {
	d1, d2, d3 := getMore()
	fmt.Printf("%s %s %d", d1.Field, d2.Field, d3.Field)
}
