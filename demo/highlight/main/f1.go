package main

import "github.com/dlsniper/u/demo/multi/m2"

func main() {
	println(message())
	println(m2.Msg())
}

type Demo interface {
	Emo(string)
}
