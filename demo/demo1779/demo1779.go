package main

type (
	de interface {
		Demo() error
	}

	dem struct{}
)

func (a dem) Demo() error {
	return error("demo")
}

func main() {
	b := dem{}
	b.Demo().Error()
}
