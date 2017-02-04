package main

import (
	"flag"
	"os"
)

type dem struct {
	Field1 string
	Field2 string
	Field3 string
}

func demo() (string, dem) {
	return "", dem{
		Field1: "a",
		Field2: "b",
		Field3: "c",
	}
}

var h bool

func init() {
	flag.BoolVar(&h, "hello", h, "heeelp")
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	a, b := demo()

	println("")
	println(a)
	println(b.Field1)
	println(b.Field2)
	println(b.Field3)
}
