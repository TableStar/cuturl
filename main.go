package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "cuturl home")
	})

	fmt.Println("server running on http://localhost:7400")
	http.ListenAndServe(":7400", nil)
}
