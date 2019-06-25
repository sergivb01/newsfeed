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

func (c *client) FetchSources() error {
	t := time.Now()
	var tempStorage struct {
		items []news.Item
		sync.Mutex
	}
	var wg sync.WaitGroup
	wg.Add(len(c.Sources))

	for _, src := range c.Sources {
		go func(src news.Source) {
			items, err := src.GetItems()
			if err != nil {
				log.Printf("error getting items for %s: %v", src.Shortname, err)
			}

			tempStorage.Lock()
			tempStorage.items = append(tempStorage.items, items...)
			tempStorage.Unlock()
			wg.Done()
		}(src)
	}
	wg.Wait()

	if !c.isEqual(tempStorage.items) {
		log.Printf("??? found != stored, saving...")
		c.Lock()
		defer c.Unlock()

		if err := c.Store.Set(tempStorage.items); err != nil {
			return fmt.Errorf("error settings items to storage: %v", err)
		}
		c.LastFetched = time.Now()
		c.NextUpdate = t.Add(c.FetchInterval)
	} else {
		log.Printf("!!! found == stored, skipping...")
	}

	time.Sleep(c.FetchInterval)

	return nil
}

func (c client) isEqual(items []news.Item) bool {
	inStore := c.Store.Get()
	if len(items) == 0 || len(inStore) == 0 {
		return false
	}

	if len(items) != len(inStore) {
		return false
	}

	for _, v := range items {
		if !containsElement(inStore, v.Link) {
			return false
		}
	}
	return true
}

func containsElement(items []news.Item, item string) bool {
	for _, v := range items {
		if v.Link == item {
			return true
		}
	}
	return false
}
