package bench2_test

import (
	"testing"
	"runtime"

	"github.com/dlsniper/u/bench/2"
)

var (
	byteEncoded []byte
	stringEncoded string
)

func BenchmarkToJSON(b *testing.B) {
	var (
		vals = make([]*bench2.Bench, b.N)
		res = []byte{}
		err error
	)

	for i:=0; i<b.N; i++ {
		vals[i] = &bench2.Bench{FieldA: "demo", FieldB: 42}
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		res, err = vals[i].ToJSON()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	byteEncoded = res
	runtime.GC()
}

func BenchmarkToJSONString(b *testing.B) {
	var (
		vals = make([]*bench2.Bench, b.N)
		res = ""
		err error
	)

	for i:=0; i<b.N; i++ {
		vals[i] = &bench2.Bench{FieldA: "demo", FieldB: 42}
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		res, err = vals[i].ToJSONString()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	stringEncoded = res
	runtime.GC()
}

func BenchmarkToJSON2(b *testing.B) {
	var (
		vals = make([]bench2.Bench, b.N)
		res = []byte{}
		err error
	)

	for i:=0; i<b.N; i++ {
		vals[i] = bench2.Bench{FieldA: "demo", FieldB: 42}
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		res, err = vals[i].ToJSON()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	byteEncoded = res
	runtime.GC()
}

func BenchmarkToJSONString2(b *testing.B) {
	var (
		vals = make([]bench2.Bench, b.N)
		res = ""
		err error
	)

	for i:=0; i<b.N; i++ {
		vals[i] = bench2.Bench{FieldA: "demo", FieldB: 42}
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		res, err = vals[i].ToJSONString2()
		if err != nil {
			b.Logf("Failed test: %d with error: %q", i, err)
			b.Fail()
		}
	}

	b.StopTimer()

	stringEncoded = res
	runtime.GC()
}
