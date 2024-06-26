package config

import (
	"encoding/json"
	"os"
)

// Config represents the structure of our configuration file
type Config struct {
	Proxies []string `json:"proxies"`
	Delay   int      `json:"delay"`
	Path    string   `json:"path"`
}

// LoadConfig reads a configuration file and decodes it into a Config struct
func LoadConfig(configPath string) (*Config, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
