package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type FetchNewBook struct{}
type GoogleResponse struct {
	Items []GoogleBook `json:"items"`
}
type GoogleBook struct {
	Kind       string     `json:"kind"`
	ID         string     `json:"id"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
}
type VolumeInfo struct {
	Title       string   `json:"title"`
	SubTitle    string   `json:"subtitle"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}

const (
	DefaultOrderBy   = "newest"
	DefaultMaxResult = "40"
)

var googleApiKey = os.Getenv("GOOGLE_API_KEY")

func (f *FetchNewBook) FetchGoogleBooks(category string) ([]GoogleBook, error) {
	rawQuery := fmt.Sprintf("key=%s&q=%s&orderBy=%s&maxResult=%s", googleApiKey, category, DefaultOrderBy, DefaultMaxResult)
	googleVolumeBuilder := &url.URL{
		Scheme:   "https",
		Host:     "www.googleapis.com",
		Path:     "/books/v1/volumes",
		RawQuery: rawQuery,
	}
	googleVolumeUrl := googleVolumeBuilder.String()
	response, err := http.Get(googleVolumeUrl)
	if err != nil {
		// Write to log
		return nil, err
	}
	defer response.Body.Close()
	var v GoogleResponse
	err = json.NewDecoder(response.Body).Decode(&v)
	books := v.Items

	if err != nil {
		// Write to log
		return nil, err
	}

	return books, nil
}
