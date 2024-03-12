// storage.go
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

func (s *Storage) Save(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.OpenFile(s.Filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = file.Write(dataBytes)
	if err != nil {
		return err
	}
	_, err = file.WriteString("\n") // Add newline for separation
	return err
}
