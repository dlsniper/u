package main

import "fmt"

func (v Vector) Negated() Vector {
	return Vector{
		X: -v.X,
		Y: -v.Y,
	}
}

type Vector struct {
	X, Y float64
}

func main() {
	a := &Vector{}
	fmt.Printf("%s", a)

	panic("omg")
}
