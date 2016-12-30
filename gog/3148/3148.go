package main

type III interface {
	FFF()
}

type EEE struct {}

func (EEE) FFF() {}

type OOO struct {
	EEE
}

func help(III) {}

func main() {
	help(OOO{})
}
