package main

import (
	handlers "cuturl/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Println("server running on http://localhost:7400")
	http.ListenAndServe(":7400", nil)
}
