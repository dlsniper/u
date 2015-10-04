package main

import "encoding/json"

type (
	di2 interface {
		De()
	}

	di interface {
		Dem() (int, di2)
	}

	d2 struct{}
	d  struct{}
)

func (a *d2) De() {}
func (a *d) Dem() (int, di2) {
	return 1, &d2{}
}

func main() {
	a := &d{}

	_, _ = json.Marshal(a)

	defer a.Dem()
}
