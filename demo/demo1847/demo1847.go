package main

func foo(_ int) {
}

func bar() {
}

func main() {
	foo(bar()) // expected "bar() used as value", actual nothing
}