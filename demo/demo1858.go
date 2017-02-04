package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%s", strconv.Quote("demo"))
	fmt.Printf("%s", "demo")
	fmt.Printf("%s", "demo")
}

func init() {
	fmt.Printf("%s", strconv.IntSize)
}
