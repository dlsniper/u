package maps_test

import (
	"strconv"
	"testing"

	"github.com/dlsniper/u/bench/maps"
)

var v string

func BenchmarkNormalGet(b *testing.B) {
	var val string
	for i := 0; i < b.N; i++ {
		val = maps.NormalGet(strconv.Itoa(i))
	}
	for i := 0; i < b.N; i++ {
		val = maps.NormalGet(strconv.Itoa(i))
	}
	v = val
	b.StopTimer()
	maps.Clear()
}

func BenchmarkPointerlGet(b *testing.B) {
	var val string
	for i := 0; i < b.N; i++ {
		val = maps.PointerGet(strconv.Itoa(i))
	}
	for i := 0; i < b.N; i++ {
		val = maps.PointerGet(strconv.Itoa(i))
	}
	v = val
	b.StopTimer()
	maps.Clear()
}
