package main

import (
	"flag"
	"fmt"
)

func main() {
	a := flag.String("demo", "d", "how to use it")
	flag.Parse()
	fmt.Printf("flag value is: %s\n", *a)
}
