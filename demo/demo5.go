package demo

type (
	demoStr struct {
		field1 int
	}

	demo struct {
		str   demoStr
		field string
	}
)

func NewDemo() *demo {
	return new(demo)
}

func d5() {
	a := NewDemo()
	a = &de

	i := 0xFFFF
	a := 1
	_ = i
	_ = a
}
