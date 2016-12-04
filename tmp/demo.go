package main

type demo int

const (
	d0 demo = iota
	d1
	d2
	d3
	d4
	// Add new value before this
	dmax
)

func main() {
	for idx := int(d0); idx < int(dmax); idx++ {
		println(idx)
	}
}
