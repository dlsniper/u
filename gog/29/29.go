package main

type A struct {

}

func (A) a1() int {
	panic("implement me")
}

func (*A) a2() int {
	panic("implement me")
}

func (*A) a3() int {
	panic("implement me")
}

func (*A) A4() int {
	panic("implement me")
}

type B interface  {
	a1() int
	a2() int
	a3() int
}

func demo(B) {}

func main() {
	a := A{}
	demo(a)
}

type myStruct struct {
	First string `de:"firstAndSecond"` // will not marshal nor unmarshal
}