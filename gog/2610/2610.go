package main

import "fmt"

type Test struct {
	Field int
}

func (t *Test) Modify() {
	t = &Test{}
}

func main() {
	d := &Test{
		Field: 1,
	}
	d.Modify()

	fmt.Printf("%#v\n", d)
}
