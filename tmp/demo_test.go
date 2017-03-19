package main_test

import "fmt"
import "testing"

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Print("hello")
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("hello")
	}
}