package main

import (
	"fmt"
	_ "github.com/dlsniper/demo1"
	_ "github.com/dlsniper/demo2"
)

type demo struct{}

func (demo) Demo() {}

func main() {
	my := demo{}
	fmt.Printf("this should be a rather long text that is made so that it can reproduce the issue GO-3496 see %s", my.Demo
}
