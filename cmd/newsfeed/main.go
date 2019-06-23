package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/sergivb01/newsfeed/platform/newsfeed"
)

func main() {
	wired := newsfeed.Source{
		Title:    "Wired",
		Homepage: "https://wired.com",
		RSS:      "https://www.wired.com/feed/rss",
		Parser: func(b []byte) ([]newsfeed.Item, error) {
			var res struct {
				Channel struct {
					LastBuildDate string `xml:"lastBuildDate"`
					Items         []struct {
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
			var items []newsfeed.Item

			for _, item := range res.Channel.Items {
				pubDate, _ := time.Parse(time.RFC1123Z, item.PublishDate)

				items = append(items, newsfeed.Item{
					Title:       item.Title,
					Link:        item.Link,
					PublishDate: pubDate,
					Description: item.Description,
					Thumbnail:   item.Thumbnail.URL,
				})
			}

			return items, nil
		},
	}

	items, err := wired.GetItems()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(items))

	b, err := json.Marshal(wired)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
