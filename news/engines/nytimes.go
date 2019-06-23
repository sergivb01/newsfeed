package engines

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/sergivb01/newsfeed/news"
)

// NYTimesSource defines wired.com source for news
var NYTimesSource = news.Source{
	Title:    "The New York Times",
	Homepage: "https://nytimes.com",
	RSS:      "https://www.nytimes.com/svc/collections/v1/publish/https://www.nytimes.com/section/technology/rss.xml",
	Parser: func(b []byte) ([]news.Item, error) {
		var res struct {
			Channel struct {
				Items []struct {
					Title       string `xml:"title"`
					Link        string `xml:"link"`
					PublishDate string `xml:"pubDate"`
					Description string `xml:"description"`
				} `xml:"item"`
			} `xml:"channel"`
		}

		if err := xml.Unmarshal(b, &res); err != nil {
			return nil, fmt.Errorf("error while unmarshaling bytes: %v", err)
		}
		var items []news.Item

		for _, item := range res.Channel.Items {
			pubDate, _ := time.Parse(time.RFC1123, item.PublishDate)

			items = append(items, news.Item{
				Title:       item.Title,
				Link:        item.Link,
				PublishDate: pubDate,
				Description: item.Description,
			})
		}
		return items, nil
	},
}
