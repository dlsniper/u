package main

import "testing"

func doStuff() {

}

func BenchmarkFoo1(b *testing.B) {
	slice := []string{"windows", "mac", "linux", "chrome"}
	myVal := "chrome"
	for i := 0; i < b.N; i++ {
		for _, val := range slice {
			if val == myVal {
				doStuff()
			}
		}
	}
}

func BenchmarkFoo2(b *testing.B) {
	slice := map[string]struct{}{"windows": {}, "mac": {}, "linux": {}, "chrome": {}}
	myVal := "chrome"
	for i := 0; i < b.N; i++ {
		if _, ok := slice[myVal]; ok {
			doStuff()
		}
	}
}

func BenchmarkFoo3(b *testing.B) {
	myVal := "chrome"
	for i := 0; i < b.N; i++ {
		switch myVal {
		case "windows":
			doStuff()
		case "mac":
			doStuff()
		case "linux":
			doStuff()
		case "chrome":
			doStuff()
		}
	}
}

func BenchmarkFoo4(b *testing.B) {
	myVal := "chrome"
	for i := 0; i < b.N; i++ {
		if myVal == "windows" {
			doStuff()
		} else if myVal == "mac" {
			doStuff()
		} else if myVal == "linux" {
			doStuff()
		} else if myVal == "chrome" {
			doStuff()
		}
	}
}

func BenchmarkFoo5(b *testing.B) {
	myVal := "chrome"
	for i := 0; i < b.N; i++ {
		if myVal == "windows" {
			doStuff()
		}
		if myVal == "mac" {
			doStuff()
		}
		if myVal == "linux" {
			doStuff()
		}
		if myVal == "chrome" {
			doStuff()
		}
	}
}
