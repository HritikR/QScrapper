// main.go
package main

import (
	"flag"
	"log"
	"nichecrawler/config"
	"nichecrawler/proxy"
	"nichecrawler/scraper"
	"nichecrawler/storage"
	"path/filepath"
)

func main() {
	startPage := flag.Int("start", 1, "Start page number")
	endPage := flag.Int("end", 1, "End page number")
	baseURL := flag.String("url", "", "Base URL with placeholder for page number")
	flag.Parse()

	// Load configuration
	cfgPath, _ := filepath.Abs("./config.json")
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize proxy manager and get all proxies
	pm := proxy.NewProxyManager(cfg.Proxies)
	allProxies := pm.GetAllProxies()

	// Initialize the storage system
	dataStorage := storage.NewStorage("./scraped_data.json")

	// Initialize the scraper
	s := scraper.NewScraper()

	// Define the processData callback
	processData := func(data map[string]interface{}) {
		if err := dataStorage.Save(data); err != nil {
			log.Printf("Failed to save data: %v", err)
		}
	}

	// Execute the scrape with all proxies available for rotation
	s.Scrape(*startPage, *endPage, *baseURL, allProxies, processData)
}
