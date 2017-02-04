package main

func demoReturn() int {
	return 1
}

type A struct{}

func main() {
	_, a := demoReturn()
	_, _ = a
	_, a, _, b := demoReturn()
	_ = b, b
	c := 1, 1, A{}
	_ = c
}
