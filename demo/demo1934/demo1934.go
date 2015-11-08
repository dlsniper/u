package demo_test

import "testing"

func BenchmarkGoroutinePerWork(b *testing.B) {
	resc := make(chan int)
	doWork := func(i int) {
		resc <- i * 2
	}
	for i := 0; i < b.N; i++ {
		go doWork(i)
		<-resc
	}
}

func BenchmarkOneWorkerGoroutine(b *testing.B) {
	workc := make(chan int)
	resc := make(chan int)
	go func() {
		for i := range workc {
			resc <- i * 2
		}
	}()
	defer close(workc)
	for i := 0; i < b.N; i++ {
		workc <- i
		<-resc
	}
}
