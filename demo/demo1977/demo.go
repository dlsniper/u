package main

type Foo struct{}

func main() {
	f := new(Foo)
	s := "string"
	s = 10
	s = f
	f = 10
	bar(f)
	println(s)

	var demo string = 10
	_ = demo
}

func bar(s string)
