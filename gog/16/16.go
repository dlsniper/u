package main

func main() {
	println(1 == 1 && !(1 == 1))
	println(!(len([]byte{}) > 0) && 1 == 1)
}
