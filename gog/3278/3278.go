package main

type Type struct{}
func (Type) TypeMethod(...int) {}
func (Type) ValueMethod(Type, ...int) {}
func F(Type, ...int) {}

func main() {
	var f func(Type, ...int)
	f = F // Ok
	f = Type{}.ValueMethod// Ok
	f = Type.TypeMethod // Cannot use Type.TypeMethod (type func (Type, int)) as type func(Type, ...int) in assignment
	_ = f
}
