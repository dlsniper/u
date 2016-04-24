package my_test

import (
	"testing"
)

func TestDemo(t *testing.T) {
	var str string
	var stri = "adjadakdjkdakadkajajakda"
	for n := 0; n < 10; n++ {
		str = str +stri
	}
	t.Logf("%     [1]d thingy %[3]d\n", 1, 2, 3)
	t.Skipf("%     [1]d thingy %[3]d\n", 1, 2, 3)
}
