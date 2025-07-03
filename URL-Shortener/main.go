package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
	ClickCount   int       `json:"click_count"`
}

var (
	urlDB = make(map[string]URL)
	dbMux = sync.Mutex{} // thread safe access
)

func generateShortURL(originalURL string) string {
	hasher := sha256.New()
	hasher.Write([]byte(originalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println("hasher: ", hash[:8])
	return hash[:8]
}
func createURL(originalURL string, expireMinutes int) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL

	dbMux.Lock()
	defer dbMux.Unlock()

	if existing, ok := urlDB[id]; ok {
		return existing.ShortURL
	}
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
		ExpiresAt:    time.Now().Add(time.Duration(expireMinutes) * time.Minute),
	}
	return shortURL
}
func getURL(id string) (URL, error) {
	dbMux.Lock()
	defer dbMux.Unlock()
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL NOT FOUND")
	}
	if time.Now().After(url.ExpiresAt) {
		delete(urlDB, id)
		return URL{}, errors.New("Url expired")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello,world!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post allowed", http.StatusMethodNotAllowed)
		return
	}
	var Data struct {
		URL         string `json:"url"`
		ExpireAfter int    `json:"expire_minutes"`
	}
	err := json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if Data.ExpireAfter <= 0 {
		Data.ExpireAfter = 60
	}
	shortURL1 := createURL(Data.URL, Data.ExpireAfter)
	response := struct {
		ShortUrl string `json:"short_url"`
	}{ShortUrl: shortURL1}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Path[len("/redirect/"):]
	url, ero := getURL(id)
	if ero != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
	}
	dbMux.Lock()
	url.ClickCount++
	urlDB[id] = url
	dbMux.Unlock()

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)

}
func listURLsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}
	dbMux.Lock()
	defer dbMux.Unlock()

	var urls []URL
	for _, u := range urlDB {
		urls = append(urls, u)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
func main() {
	//fmt.Println("Starting URL Shortener......")
	//OriginalURL := "https://www.google.co.in/"
	//generateShortURL(OriginalURL)

	//Register the handler function to handle all requests to the root URL
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)
	http.HandleFunc("/urls", listURLsHandler)

	//start the HTTP Server on port 8080
	fmt.Println("starting server on port on 3000..")
	er := http.ListenAndServe(":3000", nil)
	if er != nil {
		fmt.Println("Error on starting server", er)
	}
}
