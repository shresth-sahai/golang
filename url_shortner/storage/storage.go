package storage

import (
	"encoding/json",
	"errors",
	"fmt",
	"os",
	"sync",
	"url-shortner/models"
)

type URLStorage struct{
	urls map[string]*models.URLMapping
	mu sync.RWMutex 
	filePath string
}

func NewURLStorage(filePath string)*URLStorage{
	storage:= &URLStorage{
		urls: make(map[string]*models.URLMapping),
		filePath: filePath,
	}
	storage.LoadFromFile()
	return storage
}

func (s *URLStorage) Save(mapping *models.URLMapping) error {
	s.mu.lock()
	defer s.mu.Unlock()


	if _,exists := s.urls[mapping.ShortCode]; exists {
		return errors.New("short code already exists")
	}
	s.urls[mapping.ShortCode]=mapping

	return s.SaveToFile
}




