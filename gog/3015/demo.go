package demo

import (
	"fmt"
	"net/http"
)

func FetchStuff() {
	resp, err := http.Get("http://example.com/")
	fmt.Printf("%v\n%v", resp, err)
}

type small struct {
	f [10]int
}

type medium struct {
	d []byte
	e small
}

type large struct {
	a int
	b string
	c medium
}

func DebugDemo() {
	FetchStuff()
	a := large{
		a: 1,
		b: "demo",
		c: medium{
			d: []byte("demo"),
			e: small{
				f: [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
	}
	b := append([]large{}, a, a, a, a, a, a, a, a, a, a)
	c := a.c
	e := a.c.e
	fmt.Printf("%v\n%v\n%v\n", b, c, e)

	g := []byte("demo")
	fmt.Printf("len:\t%d\ncap:\t%d\n", len(g), cap(g))
	fmt.Printf("len:\t%d\ncap:\t%d\n", len(a.c.d), cap(a.c.d))

	h := make(chan []large, 100)
	fmt.Printf("len:\t%d\ncap:\t%d\n", len(h), cap(h))
}
