package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"

	"newsfeed.sergivos.dev/platform/newsfeed"
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

	fmt.Printf("Items: %T\nItem: %T\n", items, items[0])

	b, err := json.Marshal(items)
	if err != nil {
		fmt.Println("items arr", err)
	}

	if err := ioutil.WriteFile("temp.json", b, 0644); err != nil {
		fmt.Println(err)
	}

}
