package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "users")
	})
	http.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "lists")
	})
	http.ListenAndServe("localhost:8080", nil)
}
