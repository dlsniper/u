/*
 * Go highlight sample
 */
// Package main
package main

import (
	"errors"
	"fmt"
)

type PublicInterface interface {
	PublicFunc() int
	privateFunc() int
}
type private interface {
	PublicFunc() int
	privateFunc() int
}
type PublicStruct struct {
	PublicField  int
	privateField int
}
type privateStruct struct {
	PublicField  int
	privateField int
}
type typeSpec int
const PublicConst, privateConst = 1, 2
var PublicVar, privateVar = 1, 2

func PublicFunc() int { const LocalConst = 1; localVar := PublicVar; return localVar }
func privateFunc() (int, int) { const localConst = 2; LocalVar := privateVar; return LocalVar, PublicVar;}
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
	_, _ = PublicConst, privateConst
}

func main() { fmt.Println("demo"); variableFunc(1); c := false; d := nil; _, _ = c, d }

type demooo error

func (a demooo) demoo() error {
	return errors.New("demo")
}

func (a demooo) PublicFunc() int {
	return 1
}

func (a demooo) privateFunc() int {
	return 1
}

var a = demooo("errr")
