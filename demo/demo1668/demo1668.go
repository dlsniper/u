package main

func Foo() (int, int) {
	a, b := 0, 0
	a, b = 1, 1
	a, b /= 1, 1 // Should be error: `syntax error: unexpected op=, expecting := or = or comma`
	return a, b
}
