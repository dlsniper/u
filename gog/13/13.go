package main

func demo() {
	a := uint(1)
	print(1+1 == 1)
	print(a/a + 1)
	print(a%a*1 == 1)
	print(a << a)
	print(a >> a)
	print(a & a)
	print(a &^ a)
}

type A struct{}

func main(){
	a := A{}
	print(a.a()*a.a() == 0)
	print(a.a()+a.a() == 0)
	print(a()+a() == 0)
}
func (a A) a() {


}
