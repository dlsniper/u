package main

import (
	"fmt"
)

func change(s ...int) {
	s[0], s[1], s[2] = 4, 5, 6
}

func main() {
	mySlice := []int{1, 2, 3}
	change(mySlice[0], mySlice[1],  mySlice[2])
	fmt.Printf("%+v\n", mySlice)
	change(mySlice...)
	fmt.Printf("%+v %s\n", mySlice)

}
