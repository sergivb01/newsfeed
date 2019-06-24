package sources

import (
	"github.com/sergivb01/newsfeed/news"
)

// TheVergeSource defines wired.com source for news
var TheVergeSource = news.Source{
	Title:        "The Verge",
	Shortname:    "verge",
	Homepage:     "https://www.theverge.com/",
	RSS:          "https://www.theverge.com/rss/index.xml",
	WithChannels: false,
}
