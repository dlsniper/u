package demo1386

import "github.com/valyala/fasthttp"

type (
	// help
	Hello struct{}

	// help
	Helloi interface{}
)

// help
type Helllo struct{}

// help
type Hellloi interface{}

const (
	// help
	Helloc = 3
)

// help
const Hellloc, hellloc2 = 1, 2

var (
	// help
	Hello1, Hellow1 int
)

// help
var Hello2 int

// Helllo
func (a Helllo) Hellllo() {

}

// Demo does things  -> correct
func Demo() {
	a := 1
	_ = func(ctx *fasthttp.RequestCtx) {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	_ = a
}

//  Demo2 does other things -> incorrect
func Demo2() {

}
