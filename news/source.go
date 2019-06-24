package news

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type parserFunc func(string, []byte) ([]Item, error)

// Source defines a new news website source
type Source struct {
	Title        string `json:"title"`
	Shortname    string `json:"shortName"`
	Homepage     string `json:"homepage"`
	RSS          string `json:"rss"`
	WithChannels bool   `json:"withChannels"`
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

	parser := withoutChannels
	if s.WithChannels {
		parser = withChannels
	}

	items, err := parser(s.Shortname, b)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func withChannels(srcName string, b []byte) ([]Item, error) {
	var res struct {
		Channel struct {
			Items []struct {
				Title       string `xml:"title"`
				Link        string `xml:"link"`
				PublishDate string `xml:"pubDate"`
				Description string `xml:"description"`
				Thumbnail   struct {
					URL string `xml:"url,attr"`
				} `xml:"thumbnail"`
			} `xml:"item"`
		} `xml:"channel"`
	}

	if err := xml.Unmarshal(b, &res); err != nil {
		return nil, fmt.Errorf("error while unmarshaling bytes: %v", err)
	}
	var items []Item

	for _, rawItem := range res.Channel.Items {
		pubDate, _ := time.Parse(time.RFC1123Z, rawItem.PublishDate)
		item := Item{
			Title:       rawItem.Title,
			Link:        rawItem.Link,
			PublishDate: pubDate,
			Description: rawItem.Description,
			Thumbnail:   rawItem.Thumbnail.URL,
			Source:      srcName,
		}
		item.Clear()

		items = append(items, item)
	}
	return items, nil
}

func withoutChannels(srcName string, b []byte) ([]Item, error) {
	var res struct {
		Items []struct {
			PublishDate string `xml:"published"`
			Title       string `xml:"title"`
			Content     struct {
				Text string `xml:",chardata"`
			} `xml:"content"`
			Link struct {
				Href string `xml:"href,attr"`
			} `xml:"link"`
		} `xml:"entry"`
	}

	if err := xml.Unmarshal(b, &res); err != nil {
		return nil, fmt.Errorf("error while unmarshaling bytes: %v", err)
	}

	var items []Item
	for _, rawItem := range res.Items {
		pubDate, _ := time.Parse(time.RFC1123, rawItem.PublishDate)
		item := Item{
			Title:       rawItem.Title,
			Link:        rawItem.Link.Href,
			PublishDate: pubDate,
			Description: rawItem.Content.Text,
			Source:      srcName,
		}
		item.Clear()

		items = append(items, item)
	}

	return items, nil
}
