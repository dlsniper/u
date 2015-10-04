package main

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
