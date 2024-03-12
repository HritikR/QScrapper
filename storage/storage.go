package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type Storage struct {
	Filename string
	mu       sync.Mutex
}

func NewStorage(filename string) *Storage {
	return &Storage{
		Filename: filename,
	}
}

func (s *Storage) Save(rawData string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Ensure the file exists
	file, err := os.OpenFile(s.Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fileContent, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}

	// Initialize an empty array JSON if the file is new
	if len(fileContent) == 0 {
		fileContent = []byte("[]")
	}

	var data []interface{}
	if err := json.Unmarshal(fileContent, &data); err != nil {
		return err
	}

	var newData interface{}
	if err := json.Unmarshal([]byte(rawData), &newData); err != nil {
		return err
	}

	// Check if newData is a slice and spread it into the existing data
	if newDataSlice, ok := newData.([]interface{}); ok {
		data = append(data, newDataSlice...)
	} else {
		data = append(data, newData)
	}

	updatedContent, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename, updatedContent, 0644)
}
