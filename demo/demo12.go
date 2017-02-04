package main

type (
	Param struct {
		Key, Value string
	}

	Params []Param
)

func (p Params) ByName(name string) string {
	for i := range p {
		if p[i].Key == name {
			return p[i].Value
		}
	}
	return ""
}

func main() {
	a := Params{
		Param{"key", "value"},
	}

	println(a[0].Key, "-", a[0].Value)
	for _, p := range []Param(a) {
		println(p.Key, "-", p.Value)
	}
	println([]Param(a)[0].Key, "-", []Param(a)[0].Value)
}
