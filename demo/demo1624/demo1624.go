package main

import (
	"fmt"
)

func main() {
	m1 := new(map[string]string)
	m1["Foo"] = "Bar" // It is incorrect. But not show error in the IDEA.
	(*m1)["Bar"] = "Foo" // It is correct.
	check(m1) // It is correct.

	m2 := make(map[string]string) // It is correct.
	m2["Foo"] = "Bar" // It is correct.
	(*m2)["Bar"] = "Foo" // It is incorrect. But not show error in the IDEA.
	check(&m2)

	m3 := &make(map[string]string) // &make(...) is incorrect. But not show error in the IDEA.
	check(m3)

	m4 := &map[string]string{} // It is correct.
	m4["Foo"] = "Bar" // It is incorrect. But not show error in the IDEA.
	(*m4)["Bar"] = "Foo" // It is correct.
	check(m4) // It is correct.
}

func check(m *map[string]string) {
	fmt.Println(*m)
}
