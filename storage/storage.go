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

func (s *Storage) Save(newData interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if file exists; if not, create it and initialize with an empty array
	file, err := os.OpenFile(s.Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the current content of the file
	fileContent, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}

	// If file is empty, initialize it with an empty JSON array
	if len(fileContent) == 0 {
		fileContent = []byte("[]")
	}

	// Unmarshal the content into a slice
	var data []interface{}
	if err := json.Unmarshal(fileContent, &data); err != nil {
		return err
	}

	// Append the new data to the slice
	data = append(data, newData)

	// Marshal the updated slice back into JSON
	updatedContent, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write the updated JSON back to the file
	// Truncate the file and write from the beginning to ensure fresh content
	if err := os.WriteFile(s.Filename, updatedContent, 0644); err != nil {
		return err
	}

	return nil
}
