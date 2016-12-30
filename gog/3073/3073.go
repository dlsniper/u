package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		fmt.Fprint(w, `{"едно": "първи", "две": "втори" }`)
	})
	http.ListenAndServe(":8080", nil)
}
