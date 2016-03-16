package main

var demov = _
var (
	demov1, demov2 = _, _
)

const democ = _

const (
	democ1, democ2 = _, _
)

func (_ int) _(_ int) (_ int) {
	return 1
}

func main() {
	println(_, "hello1")
	a, b := 1, _
	_, _, _, _ = a, b, demov, democ
	_, _, _, _ = demov1, demov2, democ1, democ2
	c := 1 + _ + (_)
	println(_ + (_))
	println(c)

	select {
	case _ <- 0:  break;
	case _ <- 1:
	}

	for _, _ = range _  {
	}
	for _ = range "abc"  {
	}
}
