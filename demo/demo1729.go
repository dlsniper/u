package main

type Demo interface {
	Func1()
	func2() int
}

type demo interface {
	Func1()
}

type dem int

func (d dem) Func1() {}

func (d dem) func2() int {
	return int(d)
}

func main() {

}
