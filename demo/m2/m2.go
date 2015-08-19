package main

import (
	"os"

	"github.com/dlsniper/u/demo/m2/ctx"
)

func main() {
	cx := ctx.NewContext() // Resolves correctly to c.NewContext() :-)
	_ = cx

	println(os.Getenv("GOPATH"))
}
