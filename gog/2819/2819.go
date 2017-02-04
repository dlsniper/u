package foo

type S struct {
	X, Y int
}

func main() {
	var s S = S{}
	s = S{, Y:1}
	_ = s
}