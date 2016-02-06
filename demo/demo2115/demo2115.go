package main

type T struct {}

func (tv  T) Mv(a int) int { return 0 }

var t T

func _() {
	f2 := (T).Mv;
	f2(t, 7)
}
