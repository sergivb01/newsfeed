package engines

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/sergivb01/newsfeed/news"
)

// TheVergeSource defines wired.com source for news
var TheVergeSource = news.Source{
	Title:    "The Verge",
	Homepage: "https://www.theverge.com/",
	RSS:      "https://www.theverge.com/rss/index.xml",
	Parser: func(b []byte) ([]news.Item, error) {
		var res struct {
			Updated string `xml:"updated"`
			Items   []struct {
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
		var items []news.Item

		for _, item := range res.Items {
			pubDate, _ := time.Parse(time.RFC1123Z, item.PublishDate)

			items = append(items, news.Item{
				Title:       item.Title,
				Link:        item.Link.Href,
				PublishDate: pubDate,
				Thumbnail:   "",
			})
		}

		return items, nil
	},
}
