package sources

import (
	"github.com/sergivb01/newsfeed/news"
)

// GizmodoSource defines gizmodo.com source for news
var GizmodoSource = news.Source{
	Title:        "Gizmodo",
	Shortname:    "gizmodo",
	Homepage:     "https://es.gizmodo.com",
	RSS:          "https://es.gizmodo.com/rss",
	WithChannels: true,
}
