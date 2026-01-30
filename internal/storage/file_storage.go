package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"case.cubi.bankheist/internal/model"
)

type FileStorage struct {
	path string
	mu   sync.Mutex
}
type Storage interface {
	Load() ([]model.Pattern, error)
	Save(patterns []model.Pattern) error
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

func (s *FileStorage) Load() ([]model.Pattern, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []model.Pattern
	err = json.NewDecoder(file).Decode(&patterns)
	return patterns, err
}

func (s *FileStorage) Save(patterns []model.Pattern) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(patterns)
}
