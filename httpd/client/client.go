package client

import (
	"sync"
	"time"

	"github.com/sergivb01/newsfeed/httpd/store"
	"github.com/sergivb01/newsfeed/news"
)

// CLI stores a client
var CLI client

type client struct {
	Sources       []news.Source `json:"sources"`
	FetchInterval time.Duration `json:"fetchInterval"`

	LastFetched time.Time `json:"lastFetched"`
	NextUpdate  time.Time `json:"nextUpdate"`

	Store store.Storage `json:"-"`

	sync.Mutex `json:"-"`
}

func (c *client) UseStore(storage store.Storage) {
	c.Store = storage
}

func (c *client) GetSourceByName(name string) news.Source {
	for _, source := range c.Sources {
		if source.Shortname == name {
			return source
		}
	}
	return news.Source{}
}

// NewClient creates and sets a new client
func NewClient(fetchInterval time.Duration, sources []news.Source) {
	CLI = client{
		Sources:       sources,
		FetchInterval: fetchInterval,
		NextUpdate:    time.Now().Add(fetchInterval),
	}
}
