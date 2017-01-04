package main

import "fmt"

type (
	demoer interface {
		Read() (int, error)
	}

	demo struct{}
)

func (d *demo) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (d *demo) Close() error {
	panic("implement me")
}

func (*demo) Read() (int, error) {
	panic("implement me")
}

func (d *demo) Noop() {}

func main() {
	d := &demo{}
	a := 10
	b, err := d.Read()
	fmt.Printf("%d\n%s", a+b, err)
}
