package main

import "fmt"

type IsBlue interface {
	IsBlue()
}

type Cyan struct {
	name string
}

type Indigo struct {
	age int
}

func (Cyan) IsBlue() {}
func (Indigo) IsBlue() {}

func main() {
	fmt.Print("I'm ")
	aBlu := IsBlue(Cyan{"not"})
	switch col := aBlu.(type) {
	case Cyan, Indigo:
		if cyan, ok := col.(Cyan); ok {
			fmt.Print(cyan.name, " ")
		}
		fmt.Println("an error.")
	}
}
