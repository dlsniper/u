package main

import "github.com/google/flatbuffers/go"

type Obj struct{}

func (*Obj) Init([]byte, flatbuffers.UOffsetT) {}

//GetRootAsObj from offending code
func GetRootAsObj(buf []byte, offset flatbuffers.UOffsetT) *Obj {
	n := flatbuffers.GetUOffsetT(buf[offset:]) //cannot use buf[offset:] (type []byte) as type []byte
	x := &Obj{}
	x.Init(buf, n+offset)
	return x
}
