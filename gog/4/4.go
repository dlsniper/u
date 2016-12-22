package main

type (
	Type interface {
		Size() int64
	}
	ArrayType struct {
		Type Type
	}
)

func _(t *Type) {
	_, _ = (*t).(*ArrayType) // impossible type assertion: *ArrayType does not implement Type (missing Size method)
}
func main() {
	println("Hello, playground")
}
