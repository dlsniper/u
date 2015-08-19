/*
 * Go highlight sample
 */
// Package main
package main

import "fmt"

type (
	PublicInterface interface {
		PublicFunc() int
		// privateFunc is private
		privateFunc() int
	}
	private interface {
		PublicFunc() int
		privateFunc() int
	}
	PublicStruct struct {
		PublicField  int
		privateField int
	}
	privateStruct struct {
		PublicField  int
		privateField int
	}
	typeSpec int
)

const (
	PublicConst  = 1
	privateConst = 2
)

var (
	PublicVar  = 1
	privateVar = 2
)

func PublicFunc() int {
	const LocalConst = 1;
	localVar := PublicVar;
	fmt.Print
	return localVar
}

func privateFunc() (int, int) {
	const localConst = 2
	LocalVar := privateVar
	return LocalVar, PublicVar
}

func (ps PublicStruct) PublicFunc() int { return ps.privateField }

// privateFunc is a func like any other func
func (ps PublicStruct) privateFunc() { return ps.PublicField }

func (ps privateStruct) PublicFunc() int { return ps.privateField }

func (ps privateStruct) privateFunc() { return ps.PublicField }

func variableFunc(demo1 int) {
	demo1 = 3
	a := PublicStruct{}
	a.privateFunc()
	demo2 := 4
	if demo1, demo2 := privateFunc(); demo1 != 3 {
		_ = demo1
		_ = demo2
		return
	}
	for demo1 := range []int{1, 2, 3, 4} {
		_ = demo1
	}
	switch {
	case 1 == 2:
		demo1, demo2 := privateFunc()
		_ = demo1
		_ = demo2
	default:
		_ = demo1
	}
	b := func() int { return 1 }
	_ = b()
	_ = PublicFunc()
	_ = variableFunc(1)
	_ = demo1
	_ = demo2
}

func main() {
	fmt.Println("demo");
	variableFunc(1);
	c := false;
	d := nil;
	_, _ = c, d
}
