package main

import "github.com/shurcooL/octicons"

type demo struct{
	Field int // Deprecated
}



func main() {
	a := demo{}
	_ = a.Field

	_ = octicons.Assets
}
