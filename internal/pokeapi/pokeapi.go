package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Liamwolf56/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type LocationAreaResponse struct {
	Count    int `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
	if url == "" {
		url = baseURL + "/location-area/"
	}

	// Try cache first
	if data, found := c.cache.Get(url); found {
		fmt.Println("🔁 Using cached data")
		var result LocationAreaResponse
		err := json.Unmarshal(data, &result)
		return result, err
	}

	// Not found in cache, fetch from API
	fmt.Println("🌐 Fetching from API:", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	// Cache it
	c.cache.Add(url, body)

	var result LocationAreaResponse
	err = json.Unmarshal(body, &result)
	return result, err
}

