package main

type Demo3 struct {
	Demo2
}

func (d *Demo3) Dem3() {}

func main() {
	a := &Demo3{}
	a.Dem1()
	_ = a
}
