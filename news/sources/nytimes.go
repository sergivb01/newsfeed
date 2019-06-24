package sources

import (
	"github.com/sergivb01/newsfeed/news"
)

// NYTimesSource defines nytimes.com source for news
var NYTimesSource = news.Source{
	Title:        "The New York Times",
	Shortname:    "nytimes",
	Homepage:     "https://nytimes.com",
	RSS:          "https://www.nytimes.com/svc/collections/v1/publish/https://www.nytimes.com/section/technology/rss.xml",
	WithChannels: true,
}
