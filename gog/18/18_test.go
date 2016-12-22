package main

type demo struct {
	Field       int
	LongerField int
}

func main() {
	_ = demo{
		Field:       1,
		LongerField: 1,
	}
}
