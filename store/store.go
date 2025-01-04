package store

import "sync"

type URLStore struct {
	URLs map[string]string
	mu   sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		URLs: make(map[string]string),
	}
}

func (u *URLStore) AddURL(shortURL, longURL string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.URLs[shortURL] = longURL
}

func (u *URLStore) GetURL(shortURL string) (string, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	longURL, exists := u.URLs[shortURL]

	return longURL, exists
}
