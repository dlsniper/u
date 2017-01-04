package main

type Cmd struct {
	// The code to execute when this command is matched
	Action func()
}

func main() {
	a := Cmd{}
	a.Action()
}
