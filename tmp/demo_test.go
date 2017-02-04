package main

import (
	"fmt"
	"testing"
)

var chars = []byte{'{', '[', '"', '-', '0', 't', 'f', 'n'}
var charlen = len(chars)

func BenchmarkSwitch(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		switch chars[i%charlen] {
		case '{':
			m++
		case '[':
			m++
		case '"':
			m++
		case '-':
			m++
		case '0':
			m++
		case 't':
			m++
		case 'f':
			m++
		case 'n':
			m++
		}
	}
	fmt.Println(m)
}

func BenchmarkIf(b *testing.B) {
	m := 0
	for i := 0; i < b.N; i++ {
		c := chars[i%charlen]
		if c == '{' {
			m++
		} else if c == '[' {
			m++
		} else if c == '"' {
			m++
		} else if c == '-' {
			m++
		} else if c == '0' {
			m++
		} else if c == 't' {
			m++
		} else if c == 'f' {
			m++
		} else if c == 'n' {
			m++
		}
	}
	fmt.Println(m)
}

// $ go test -bench=Benchmark*
// BenchmarkValid-8    	100000000	        11.7 ns/op	       0 B/op	       0 allocs/op
// 1
// BenchmarkSwitch-8   	100
// 10000
// 1000000
// 100000000
// 200000000
// 200000000	         8.39 ns/op	       0 B/op	       0 allocs/op
// 1
// BenchmarkIf-8       	100
// 10000
// 1000000
// 100000000
// 200000000
// 200000000	         9.37 ns/op	       0 B/op	       0 allocs/op
// PASS
