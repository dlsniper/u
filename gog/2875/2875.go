package t

type A struct {

}

func (A) a1() int {
	panic("implement me")
}

func (*A) a2() int {
	panic("implement me")
}

type B interface  {
	a1() int
	a2() int
	a3() int
}
