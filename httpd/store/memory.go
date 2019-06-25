package store

import (
	"github.com/sergivb01/newsfeed/news"
)

type MemoryStore struct {
	Items []news.Item
}

// Get gets items from the in memory database
func (s *MemoryStore) Get() []news.Item {
	return s.Items
}

// Set sets what items are in the storage
func (s *MemoryStore) Set(items []news.Item) error {
	s.Items = items
	return nil
}

// InMemory makes a newsfeed in memory
func InMemory() *MemoryStore {
	return &MemoryStore{
		Items: []news.Item{},
	}
}
