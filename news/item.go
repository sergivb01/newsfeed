package news

import (
	"time"
)

const maxLength = 300

// Item defines a single "news"
type Item struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	PublishDate time.Time `json:"published"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	Source      string    `json:"source"`
}

// Clear clears the description
// TODO: Automatically extract image from description and set it as thumbnail
func (item *Item) Clear() {
	num := maxLength
	if len(item.Description) > num {
		if num > 3 {
			num -= 3
		}
		item.Description = item.Description[0:num] + "..."
	}
}
