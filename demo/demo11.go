package main

type type1 struct {
	var1    int
	var2    int
	string1 int
}

func main() {
	struct1 := type1{var1: 1}
	println(struct1.var2)
	println(struct1)
}
