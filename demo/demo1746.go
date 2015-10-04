package main

import (
	"fmt"

	"github.com/axgle/mahonia"
)

func main() {
	dc := mahonia.NewDecoder("GB18030")
	fmt.Println(dc.ConvertString("abc123"))
}

type Demo func(int)

func newDemo() Demo {
	return func(param int) {
		println(param)
	}
}

func (d Demo) Func() Demo {
	println("Func on Demo")
	return func(param int) {
		println("Demo from Func on Demo")
	}
}

func main() {
	a := newDemo()
	a.Func()
}
