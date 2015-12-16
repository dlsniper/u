package main

import (
	"m2"
	"fmt"
)

func main() {
	//println(message())
	fmt.Println(m2.Msg())
}

type Demo interface {
	Emo(string)
}
