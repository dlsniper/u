package main

import (
	"fmt"
	"log"
	"net/http"
)

type demo struct {
	hello string `json:"hello"`
	jello string `json:"jello"`
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
	d := demo{
		hello: "hello",
		jello: "jello",
	}
	_ = d
	handlers := http.NewServeMux()
	handlers.HandleFunc("/", handler)
	server := &http.Server{Addr: ":8889", Handler: handlers}
	log.Fatal(server.ListenAndServe())
}
