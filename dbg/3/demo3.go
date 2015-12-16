package main

import (
	"fmt"
	"log"
	"net/http"
	foo "net/http"
)

type demo struct {
	hell string `json:"hello"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	a := 1
	b := 2
	if a == b {
		fmt.Printf("demo %d %d", a, b)
	}
	handlers := http.NewServeMux()
	handlers.HandleFunc("/", handler)
	server := &http.Server{Addr: ":8889", Handler: handlers}
	log.Fatal(server.ListenAndServe())
	foo.
}
