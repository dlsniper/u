package demo

import (
	"fmt"
	"net/http"
)

func FetchStuff() {
	resp, err := http.Get("http://example.com/")
	fmt.Printf("%v\n%v", resp, err)
}
