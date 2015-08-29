package main

import (
	"log"
	"time"
)

const world = `World`

func main() {
	time.After(1 * time.Second)

	log.Println("Hello " + world + "!")
	for {
		<-time.After(2 * time.Second)
		log.Println("Hello " + world + "!")
		log.Panicf("omfg")
	}
}