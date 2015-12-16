package main

func foo(_, _ int) {
}

func bar() {
}

func main() {
	foo(1, bar()) // expected "bar() used as value", actual nothing
}
