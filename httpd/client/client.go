package client

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/sergivb01/newsfeed/httpd/store"
	"github.com/sergivb01/newsfeed/news"
)

// CLI stores a client
var CLI client

type client struct {
	Sources       []news.Source
	FetchInterval time.Duration

	LastFetched time.Time
	NextUpdate  time.Time

	Store store.Storage

	sync.RWMutex
}

func (c *client) UseStore(storage store.Storage) {
	c.Store = storage
}

// * This is being used on the frontend to display "Wired" on each card
func (c *client) GetSourceByName(name string) news.Source {
	for _, source := range c.Sources {
		if source.Title == name {
			return source
		}
	}
	return news.Source{}
}

// NewClient creates a new client from the yaml configuration file
func NewClient(filePath string) error {
	c, err := loadConfig(filePath)
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	interval, err := time.ParseDuration(c.FetchInterval)
	if err != nil {
		return fmt.Errorf("error while parsing duration (%s): %v", c.FetchInterval, err)
	}

	CLI = client{
		Sources:       c.Sources,
		FetchInterval: interval,
	}

	return nil
}

// FetchSources loops around each source and fetches the news. It aggregates all the news
func (c *client) FetchSources() error {
	var tempStorage struct {
		items []news.Item
		sync.Mutex
	}

	var wg sync.WaitGroup

	for _, src := range c.Sources {
		wg.Add(1)
		go func(src news.Source) {
			items, err := src.GetItems()

			// not returning an error because I won't care if a single source is down. Not to be bothered to implement channels
			if err != nil {
				log.Printf("error getting items for %s: %v", src.Title, err)
			}

			tempStorage.Lock()
			tempStorage.items = append(tempStorage.items, items...)
			tempStorage.Unlock()
			wg.Done()
		}(src)
	}
	wg.Wait()

	c.Lock()
	defer c.Unlock()

	if err := c.Store.Set(tempStorage.items); err != nil {
		return fmt.Errorf("error settings items to storage: %v", err)
	}

	c.LastFetched = time.Now()
	c.NextUpdate = time.Now().Add(c.FetchInterval)

	return nil
}
