package main

import (
	handlers "cuturl/handlers"
	"fmt"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (for dev)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle OPTIONS (pre-flight request)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass request to the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {

	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/redir/", handlers.RedirectHandler)

	http.Handle("/", enableCORS(http.DefaultServeMux))

	fmt.Println("server running on http://localhost:7400")
	http.ListenAndServe(":7400", nil)
}
