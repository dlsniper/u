package main

import "fmt"

type type1 struct {
	var1    int
	Var3    int
	var2    int
	string1 int
}

func main() {
	struct1 := type1{var1: 1, Var3: 1} // var1 is FIELD_NAME here
	fmt.Printf("%v", struct1)          // var2 is REFERENCE_EXPRESSION here
}
