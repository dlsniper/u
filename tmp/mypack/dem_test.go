package mypack_test

import (
	"testing"
)

func TestDemo(t *testing.T) {
	var str string
	var stri = "adjadakdjkdakadkajajakda"
	for n := 0; n < 10; n++ {
		str = str+stri
	}
}