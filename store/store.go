package store

import (
	"sync"
)

type URLStore struct {
	URLs map[string]string
	mu   sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		URLs: make(map[string]string),
	}
}

func (u *URLStore) AddURL(shortURL, longURL string) string {
	u.mu.Lock()
	defer u.mu.Unlock()

	for existingShortURL, existingLongURL := range u.URLs {
		if existingLongURL == longURL {
			return existingShortURL
		}
	}

	u.URLs[shortURL] = longURL
	return shortURL
}

func (u *URLStore) GetURL(shortURL string) (string, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	longURL, exists := u.URLs[shortURL]

	return longURL, exists
}

func (u *URLStore) GetAllURLs() map[string]string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	// Create a copy to avoid external modification
	urlsCopy := make(map[string]string)
	for shortURL, longURL := range u.URLs {
		urlsCopy[shortURL] = longURL
	}
	return urlsCopy
}
