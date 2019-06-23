package news

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type parserFunc func([]byte) ([]Item, error)

// Source defines a new news website source
type Source struct {
	Title    string `json:"title"`
	Homepage string `json:"homepage"`
	RSS      string `json:"rss"`

	LastUpdated time.Time `json:"lastUpdated"`
	LastFetched time.Time `json:"lastFetched"`

	Parser parserFunc `json:"-"`
}

// Fetch returns an array of bytes from the fetched RSS feed or an error
func (s Source) fetch() ([]byte, error) {
	resp, err := http.Get(s.RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}

// GetItems fetches the RSS feed, parses the news and returns an slice of items
func (s *Source) GetItems() ([]Item, error) {
	b, err := s.fetch()
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch items for %s: %v", s.Title, err)
	}

	items, err := s.Parser(b)
	if err != nil {
		return nil, err
	}

	s.LastFetched = time.Now()

	return items, nil
}
