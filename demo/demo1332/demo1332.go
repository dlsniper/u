package demo1332

type demoType int

func (d demoType) hello() {
	_ = d
}

const con = 1
var a int

func demoFunc(b int) {
	c := []int{}
	d := c
	a = con
	a = b
	_ = a
	_ = len(c)
	_ = c[0]
	_ = d
	demoFunc(1)
	_ = demoFunc
}
