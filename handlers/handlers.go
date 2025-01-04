package handlers

import (
	store "cuturl/store"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"path"
)

var urlStore = store.NewURLStore()

var baseURL = os.Getenv("BASE_URL")

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request Method", http.StatusMethodNotAllowed)
		return
	}
	longURL := r.FormValue("longURL")
	if longURL == "" {
		http.Error(w, "Missing long URL", http.StatusBadRequest)
		return
	}

	generatedShortURL := generateShortURL()

	shortURL := urlStore.AddURL(generatedShortURL, longURL)

	// fmt.Fprintf(w, "Shortened URL: http://localhost:7400/%s", shortURL)

	if baseURL == "" {
		baseURL = "http://localhost:7400"
	}

	response := map[string]string{
		"shortened_url": baseURL + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := path.Base(r.URL.Path)
	if shortURL == "/" || shortURL == "" {
		http.Error(w, "Invalides URL", http.StatusBadRequest)
		return
	}

}

func generateShortURL() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortURL)
}
