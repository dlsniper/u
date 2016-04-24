package my

import "fmt"

type test struct {
	field1 string
}

func newTest() test {
	for {
		println("")
	}

	for i:=0; i<10; i++ {

	}

	for range []byte{} {
	}

	for i:=0; i<10; i++ {
		println("hello")
	}

	for {
	}

	return test{field1: ""}
}

func _() {
	fmt.Errorf("%d %d %#[1]x %#x %2.f %d %2.2f %.f %.3f %[9]*.[2]*[3]f %d %f %#[1]x %#x %[2]d %v % d",
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	)
}
