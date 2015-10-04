package bench6

import (
	"io"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)

	return mux
}

func Server() {
	mux := Mux()
	http.ListenAndServe(":8000", mux)
}