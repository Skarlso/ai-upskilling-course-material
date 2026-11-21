package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"example.com/pantry/internal/item"
)

// Load reads a store from a JSON file. A file that does not exist yet is an
// empty pantry, not an error.
func Load(path string) (*Store, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return New(), nil
	}
	if err != nil {
		return nil, err
	}
	var items []item.Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("store: reading %s: %w", path, err)
	}
	s := New()
	for _, it := range items {
		if err := s.Add(it); err != nil {
			return nil, fmt.Errorf("store: reading %s: %w", path, err)
		}
	}
	return s, nil
}

// Save writes the store to a JSON file, sorted by name.
func (s *Store) Save(path string) error {
	data, err := json.MarshalIndent(s.List(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}
