package client

import (
	"sync"
	"time"

	"github.com/sergivb01/newsfeed/httpd/store"
	"github.com/sergivb01/newsfeed/news"
)

var CLI client

type client struct {
	Sources       map[string]news.Source `json:"sources"`
	FetchInterval time.Duration          `json:"fetchInterval"`

	LastFetched time.Time `json:"lastFetched"`
	NextUpdate  time.Time `json:"nextUpdate"`

	Store store.Storage `json:"-"`

	sync.Mutex `json:"-"`
}

func (c *client) UseStore(storage store.Storage) {
	c.Store = storage
}

// NewClient creates and sets a new client
func NewClient(fetchInterval time.Duration, sources map[string]news.Source) {
	CLI = client{
		Sources:       sources,
		FetchInterval: fetchInterval,
		NextUpdate:    time.Now().Add(fetchInterval),
	}
}
