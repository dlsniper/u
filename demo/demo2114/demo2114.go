package main

func hello() func() string {
return nil
}

func main() {
	hello()() // This should have a warning as it will crash when hello() returns nil
}