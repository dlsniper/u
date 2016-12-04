package mypack

type S interface{}
type SNope interface{
	Nope(S)
}
