package main

type Demo3 struct {
	Demo2
}

func (d *Demo3) Demo3() {}

func main() {
	a := Demo3{}

	_ = a
}
