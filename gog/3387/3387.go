package main

type demo func(format string, v ...interface{})

func noop(string, ...interface{}) {}

func something(a demo) {
	_ = a
}

func main() {
	something(noop)
}
