package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Scraper struct {
	Client *http.Client
}

func NewScraper() *Scraper {
	return &Scraper{
		// Initialize the HTTP client with a default timeout
		Client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *Scraper) Scrape(startPage, endPage int, baseURL string, proxies []string, processData func(data map[string]interface{})) {
	// If no proxies are provided, add an empty string to the slice to run the loop once without a proxy
	if len(proxies) == 0 {
		proxies = append(proxies, "")
	}

	for _, proxyURL := range proxies {
		if proxyURL != "" {
			log.Printf("Using proxy: %s", proxyURL)
			proxy, err := url.Parse(proxyURL)
			if err != nil {
				log.Printf("Error parsing proxy URL: %v", err)
				continue // Skip to the next proxy if the current one is invalid
			}
			s.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
		} else {
			log.Printf("Not using any proxy")
			s.Client.Transport = &http.Transport{} // Reset to default transport without proxy
		}

		for page := startPage; page <= endPage; page++ {
			pageURL := strings.Replace(baseURL, "{page}", fmt.Sprintf("%d", page), 1)
			log.Printf("Scraping page %d", page)
			req, err := http.NewRequest("GET", pageURL, nil)

			if err != nil {
				log.Printf("Error creating request: %v", err)
				continue
			}

			resp, err := s.Client.Do(req)
			if err != nil {
				log.Printf("Error making request: %v", err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 && resp.StatusCode != 201 {
				log.Printf("Received status code %d, switching proxy...", resp.StatusCode)
				break // Exit the current proxy loop to try the next proxy
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Error reading response body: %v", err)
				continue
			}

			// log.Printf("Received response: %s", string(body))

			var result map[string]interface{}
			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("Error unmarshalling response: %v", err)
				continue
			}

			processData(result)

			log.Printf("Sleeping for 2 seconds before next request...")
			// Add a delay of 2 seconds before the next request
			time.Sleep(2 * time.Second)
		}
	}
}
