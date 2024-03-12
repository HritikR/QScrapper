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
	"time"
)

func main() {
	startPage := flag.Int("start", 1, "Start page number")
	endPage := flag.Int("end", 1, "End page number")
	baseURL := flag.String("url", "", "Base URL with placeholder for page number")
	output := flag.String("out", "output.json", "Output file path")
	flag.Parse()

	// Load configuration
	cfgPath, _ := filepath.Abs("./config.json")
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Print("Configuration not found or invalid configuration file. Resuming without proxiesc.")
		cfg = &config.Config{Proxies: []string{}, Delay: 5 * time.Second}
	}

	// Initialize proxy manager and get all proxies
	pm := proxy.NewProxyManager(cfg.Proxies)
	allProxies := pm.GetAllProxies()

	// Initialize the storage system
	storageFile := "output.json"
	if *output != "" {
		storageFile = *output
	}
	dataStorage := storage.NewStorage(storageFile)

	// Initialize the scraper
	s := scraper.NewScraper()

	// Define the processData callback
	processData := func(data string) {
		if err := dataStorage.Save(data); err != nil {
			log.Printf("Failed to save data: %v", err)
		}
	}

	// Execute the scrape with all proxies available for rotation
	s.Scrape(*startPage, *endPage, *baseURL, cfg.Delay, cfg.Path, allProxies, processData)
}
