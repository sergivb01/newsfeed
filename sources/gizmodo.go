package sources

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"time"

	"github.com/sergivb01/newsfeed/news"
)

var regx = regexp.MustCompile("<item>(.*?)</item>")

// GizmodoSource defines gizmodo.com source for news
var GizmodoSource = news.Source{
	Title:    "Gizmodo",
	Homepage: "https://es.gizmodo.com",
	RSS:      "https://es.gizmodo.com/rss",
	Parser: func(b []byte) ([]news.Item, error) {
		s := string(b)
		xmlItems := regx.FindAllString(s, -1)

		var items []news.Item
		for _, xmlItem := range xmlItems {
			var res struct {
				Title       string `xml:"title"`
				Link        string `xml:"link"`
				Description string `xml:"description"`
				PubDate     string `xml:"pubDate"`
			}
			if err := xml.Unmarshal([]byte(xmlItem), &res); err != nil {
				fmt.Printf("error while unmarshaling bytes: %v", err)
				continue
			}

			pubDate, _ := time.Parse(time.RFC1123, res.PubDate)

			items = append(items, news.Item{
				Title:       res.Title,
				Link:        res.Link,
				Description: cutText(res.Description),
				PublishDate: pubDate,
				Source:      "gizmodo",
			})

		}
		return items, nil
	},
}
