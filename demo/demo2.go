package main

import "fmt"

// d is a function like all other functions
func d(
	demo interface{},
	err error,
) (
	[]interface{},
	error,
) {
	return []interface{}{}, err
}

func d2(s string, args ...interface{}) (interface{}, error) {
	x := make(map[string]int)
	x["foo"] = 1

	if _, ok := x["foo"]; !ok {
		fmt.Println("Missing key 'foo'")
	}

	return nil, nil
}

// d3 func comment
func d3(s string) string {
	return s
}

func main() {
	fmt.
		_, _ = d(d2("demo", d))
}
