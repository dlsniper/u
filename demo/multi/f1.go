package main

import "github.com/dlsniper/u/demo/multi/Godeps/_workspace/m2"

func main() {
	println(message())
	println(m2.Msg())
}

type Demo interface {
	Emo(string)
}
