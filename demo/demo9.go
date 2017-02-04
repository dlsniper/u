package main

func main() {
	type demo struct {
		Field1, Field2 int
	}

	var (
		Field1 int
		Field2 int
	)

	Field1 = 1
	Field2 = 2

	_, _ = Field1, Field2
}
