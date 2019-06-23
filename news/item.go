package news

import (
	"time"
)

// Item defines a single "news"
type Item struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	PublishDate time.Time `json:"published"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
}
