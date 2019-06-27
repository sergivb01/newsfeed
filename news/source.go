package news

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeLayouts = []string{
	time.RFC1123Z,
	time.RFC1123,
	time.RFC3339,
}

// Source defines a new news website source
type Source struct {
	Title        string `json:"title" yaml:"title"`
	Homepage     string `json:"homepage" yaml:"homepage"`
	RSS          string `json:"rss" yaml:"rss"`
	WithChannels bool   `json:"withChannels" yaml:"withChannels"`
}

// fetch returns an array of bytes from the fetched RSS feed or an error
func (s Source) fetch() ([]byte, error) {
	resp, err := http.Get(s.RSS)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// GetItems fetches the RSS feed, parses the news and returns an slice of items
func (s *Source) GetItems() ([]Item, error) {
	b, err := s.fetch()
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch items for %s: %v", s.Title, err)
	}

	parser := parseWithoutChannels
	if s.WithChannels {
		parser = parseWithChannels
	}

	return parser(s.Title, b)
}

func parseTime(rawTime string) time.Time {
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, rawTime); err == nil {
			return t
		}
	}
	var res time.Time
	return res
}

func parseWithChannels(srcName string, b []byte) ([]Item, error) {
	var res withChannels
	if err := xml.Unmarshal(b, &res); err != nil {
		return nil, fmt.Errorf("error while unmarshaling bytes: %v", err)
	}

	var items []Item
	for _, rawItem := range res.Channel.Items {
		item := Item{
			Title:       rawItem.Title,
			Link:        rawItem.Link,
			PublishDate: parseTime(rawItem.PublishDate),
			Description: rawItem.Description,
			Source:      srcName,
		}
		item.Clear()

		items = append(items, item)
	}
	return items, nil
}

func parseWithoutChannels(srcName string, b []byte) ([]Item, error) {
	var res withoutChannels
	if err := xml.Unmarshal(b, &res); err != nil {
		return nil, fmt.Errorf("error while unmarshaling bytes: %v", err)
	}

	var items []Item
	for _, rawItem := range res.Items {
		item := Item{
			Title:       rawItem.Title,
			Link:        rawItem.Link.Href,
			PublishDate: parseTime(rawItem.PublishDate),
			Description: rawItem.Content.Text,
			Source:      srcName,
		}
		item.Clear()

		items = append(items, item)
	}

	return items, nil
}
