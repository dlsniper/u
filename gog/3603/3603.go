package main

import (
"fmt"
"strconv"
)

func demo(v bool) {
	a := "string"
	b := 2

	switch b {
	case 0:
		v:=false
		fmt.Println(v)
	case 1:
		fmt.Println(a)
	case 2:
		v, err := strconv.ParseFloat(a, 64)
		fmt.Printf("%T %v %s \n", v, v, v)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	demo(false)
}
