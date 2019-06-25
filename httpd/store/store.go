package store

import (
	"github.com/sergivb01/newsfeed/news"
)

// Storage defines the abstraction of an storage system to save
// different Items from different news sources
type Storage interface {
	// Get returns an array of the stored items
	Get() []news.Item
	// Set sets the items in the storage
	Set([]news.Item) error
}
