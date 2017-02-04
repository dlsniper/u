package main

import "os"

func main() {
	port := os.Getenv("PORT")
	println(port)
}
