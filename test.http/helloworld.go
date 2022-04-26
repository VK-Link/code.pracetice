package main

import (
	"fmt"
	"net/http"
)

func indexHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", indexHandler2)
	http.ListenAndServe(":8000", nil)
}
