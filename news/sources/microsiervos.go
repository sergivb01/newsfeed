package sources

import (
	"github.com/sergivb01/newsfeed/news"
)

// MicrosiervosSource defines gizmodo.com source for news
var MicrosiervosSource = news.Source{
	Title:        "Microsiervos",
	Shortname:    "microsiervos",
	Homepage:     "https://www.microsiervos.com",
	RSS:          "https://www.microsiervos.com/index.xml",
	WithChannels: true,
}
