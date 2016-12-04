package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	out := struct {
		J string `json:"j"`
	}{}
	json.NewDecoder(io.TeeReader(req.Body, os.Stdout)).Decode(&out)
	log.Printf("got: %#v\n", out)
	io.WriteString(w, "Hello world!")
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)

	return mux
}

func main() {
	mux := Mux()
	http.ListenAndServe(":8000", mux)
}
