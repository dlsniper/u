package main

import "net/http"

func main() {
	http.ListenAndServe(":8081", stdlibSrv{})
}

type stdlibSrv struct{}

func (s stdlibSrv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{ "greeting" : "Hello world!" }`))
}
