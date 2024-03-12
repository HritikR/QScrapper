package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Scraper struct {
	Client *http.Client
}

func NewScraper() *Scraper {
	return &Scraper{
		Client: &http.Client{},
	}
}

func (s *Scraper) Scrape(startPage, endPage int, baseURL string, proxies []string, processData func(data map[string]interface{})) {
	for _, proxyURL := range proxies {
		log.Printf("Using proxy: %s", proxyURL)
		proxy, _ := url.Parse(proxyURL)
		s.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}

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

			var result map[string]interface{}
			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("Error unmarshalling response: %v", err)
				continue
			}

			processData(result)
		}
	}
}
