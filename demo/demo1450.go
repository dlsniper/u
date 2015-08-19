package main

import "fmt"

func myNew(v bool) { // flag this as overriding the new function from builtin language functions
	fmt.Printf("Got %v\n", v)
}

func main() {
	var string = 3
	int, myOtherVar := true, 3 // Unused variable 'myVar'
	myNew(int)                 // Unresolved type 'myVar'
	_, _ = string, myOtherVar
}
