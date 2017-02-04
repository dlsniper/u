package main

import (
	"fmt"
	"strconv"
)

func main() {
	ShouldAllowParameterShadowing("param")
}

func ShouldAllowParameterShadowing(i string) {
	fmt.Println(i)

	for i := 0; i < 3; i++ { // Inspection error "No new variables on left side of :="
		fmt.Println(strconv.Itoa(i))
	}
}
