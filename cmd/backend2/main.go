package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from server on port 4000")
	})

	fmt.Printf("Starting test server on port 4000")
	http.ListenAndServe(":4000", nil)
}
