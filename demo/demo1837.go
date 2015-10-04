package main

type TestStruct struct {
	SomeId int
}

type DataSlice []*TestStruct

func ErrorCode() {
	data := DataSlice{}
	for _, element := range data {
		if element.SomeId > 20 {
			println("some text")
		}
	}
}
