package main

import (
	"fmt"
)

func main() {
	funcName()
}
func funcName() {
	var sample rune = '\u4e16'
	var sample2 rune = '\u4e17'
	fmt.Println(sample, sample2)
}
