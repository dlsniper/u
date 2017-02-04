package main

func main() {
	var s string = "hello"
	s[0] = 0 // error
	for s[0] = range "asdf" {
	}
}
