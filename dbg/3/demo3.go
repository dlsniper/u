package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {

	handlers := http.NewServeMux()
	handlers.HandleFunc("/", handler)
	server := &http.Server{Addr: ":8889", Handler: handlers}

	server.ListenAndServe()
}
