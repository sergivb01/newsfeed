package sources

import (
	"github.com/sergivb01/newsfeed/news"
)

// WiredSource defines wired.com source for news
var WiredSource = news.Source{
	Title:        "Wired",
	Shortname:    "wired",
	Homepage:     "https://wired.com",
	RSS:          "https://www.wired.com/feed/rss",
	WithChannels: true,
}
