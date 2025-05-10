package main

import (
	"fmt"
	"net/http"
)

// Сервер для тестирования работы балансировщика
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Response from server on port 8000")
	})

	fmt.Printf("Starting test server on port 8000")
	http.ListenAndServe(":8000", nil)
}
