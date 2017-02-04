package demo2069

type AStruct struct {
	Value string
}

type AFunc func() []AStruct

func Bug() {
	var f AFunc

	structs := f()
	for idx, s := range structs {
		_ = structs[idx].Value
		_ = s.Value
	}
}
