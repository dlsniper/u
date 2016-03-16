package main

import "reflect"

func demo(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	_, _ = t, v
}

func main() {
	demo(nil)
}
