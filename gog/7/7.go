package main

import "io"

type demo struct{}

func (d demo) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (d demo) Read(p []byte) (n int, err error) {
	panic("implement me")
}

/*func (d demo) Write(p []byte) (n int, err error) {
	panic("implement me")
}*/

func main() {
	var a io.ReadWriter = demo{}
	_ = a
}
