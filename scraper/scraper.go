package scraper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"qscrapper/parser"
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

func (s *Scraper) Scrape(startPage, endPage int, baseURL string, delay time.Duration, dataPath string, proxies []string, processData func(string)) {
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
			log.Printf("Scraping page %s", pageURL)
			req, err := http.NewRequest("GET", pageURL, nil)

			// Add a user agent header to the request
			req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")

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

			// Parse JSON
			data := parser.ParseJSONData(string(body), dataPath)

			processData(data)
			log.Printf("Processed page %d", page)

			if page == endPage {
				log.Printf("Finished scraping all pages")
				break
			}

			log.Printf("Sleeping for %v seconds before next request...", delay.Seconds())
			// Add a delay before the next request
			time.Sleep(delay)
		}
	}
}
