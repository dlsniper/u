package main

type demo struct{}

func f(d demo) {}

var (
	c []struct{}
)

func main() {
	a := demo{}
	f(a) // type declaration works
	b := make(chan int)
	_ = b // type declaration does not work
	_ = c // type declaration does not work
}
