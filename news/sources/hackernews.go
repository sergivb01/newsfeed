package sources

import (
	"github.com/sergivb01/newsfeed/news"
)

// HackerNewsSource defines news.ycombinator.com source for news
var HackerNewsSource = news.Source{
	Title:        "Hacker News",
	Shortname:    "hackernews",
	Homepage:     "https://news.ycombinator.com",
	RSS:          "https://hnrss.org/frontpage",
	WithChannels: false,
}
