package main

import "io/ioutil"

func main() {
	b := ""
	_ = "a" + b
	ioutil.ReadAll()
	c := nil
	_ = c
}
