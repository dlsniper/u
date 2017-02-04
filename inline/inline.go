package main

import (
	"errors"
)

type (
	context struct {
		field1 string
		field2 int
	}

	handlerFunc func(*context) error
)

var (
	errWrongField1Value = errors.New("wrong field1 value")
	errWrongField2Value = errors.New("wrong field2 value")

	handlers = map[string][]handlerFunc{
		"index": []handlerFunc{firstFunc, secondFunc},
	}
)

func (ctx *context) funcA() string {
	return ctx.field1
}

func (ctx *context) funcB() int {
	return ctx.field2
}

func firstFunc(ctx *context) error {
	if ctx.field1 != "field1" {
		return errWrongField1Value
	}

	// Perform actions here

	return nil
}

func secondFunc(ctx *context) error {
	if ctx.field2 != 2 {
		return errWrongField2Value
	}

	// Perform more actions here

	return nil
}

func newContext() *context {
	return &context{
		field1: "field1",
		field2: 2,
	}
}

func main() {
	path := "index"
	exec := handlers[path]
	ctx := newContext()
	for idx := range exec {
		if err := exec[idx](ctx); err != nil {
			println("huston we have a problem")
		}
	}
	println("all good")
}
