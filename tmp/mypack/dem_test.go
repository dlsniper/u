package mypack_test

import "testing"

/*

This crashes my machine when running: go test -bench=. -benchtime 3m

*/

func Benchmark_SpawnRoutines(b *testing.B) {
	var a = 100
	var doSomething = func(a int) {
		a = 1000
	}
	for i := 0; i < b.N; i++ {
		go doSomething(a)
	}
}
