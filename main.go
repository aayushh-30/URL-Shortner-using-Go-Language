package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// URL struct represents a shortened URL with metadata
type URL struct {
	Id           string    `json:"id"`
	OriginalURL  string    `json:"original_URL"`
	ShortURL     string    `json:"short_URL"`
	CreationDate time.Time `json:"creation_Date"`
}

// Global map to store shortened URLs
var URL_Map = make(map[string]URL)
// id ---> {	id : 
// 				OriginalURL : 
// 				ShortURL :
// 				CreationDate
// 			 }

// generateShortURL generates an 8-character hash for the given URL
func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8] // Use the first 8 characters as the short URL
}

// CreateURL generates a short URL and stores it in the map
func CreateURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	URL_Map[id] = URL{
		Id:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// GetOriginalURL retrieves the original URL from the map using the short URL
func GetOriginalURL(shortURL string) (URL, error) {
	url, exists := URL_Map[shortURL]
	if !exists {
		return URL{}, errors.New("URL NOT FOUND")
	}
	return url, nil
}

// handler is the root endpoint that indicates the server status
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is Running")
}

// ShortURLHandler handles URL shortening requests
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create and store the shortened URL
	shortURL := CreateURL(data.URL)

	// Return the shortened URL as JSON response
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// redirectURLHandler redirects to the original URL based on the short URL
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL ID from the path
	id := r.URL.Path[len("/redirect/") : len(r.URL.Path)]

	// Retrieve the original URL from the map
	url, err := GetOriginalURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	fmt.Println("Server is Ready on http://localhost:8000")

	// Route handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the server
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
