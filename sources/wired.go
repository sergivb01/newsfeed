package sources

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/sergivb01/newsfeed/news"
)

// WiredSource defines wired.com source for news
var WiredSource = news.Source{
	Title:    "Wired",
	Homepage: "https://wired.com",
	RSS:      "https://www.wired.com/feed/rss",
	Parser: func(b []byte) ([]news.Item, error) {
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
		var items []news.Item

		for _, item := range res.Channel.Items {
			pubDate, _ := time.Parse(time.RFC1123Z, item.PublishDate)

			items = append(items, news.Item{
				Title:       item.Title,
				Link:        item.Link,
				PublishDate: pubDate,
				Description: cutText(item.Description),
				Thumbnail:   item.Thumbnail.URL,
				Source:      "wired",
			})
		}
		return items, nil
	},
}
