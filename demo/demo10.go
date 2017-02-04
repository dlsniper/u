package demo

type CompleteMe struct {
	Name    string
	NotName int
	Truthy  bool
}

var completeMe = CompleteMe{}

func inFunc() {
	completed := CompleteMe{}
}
