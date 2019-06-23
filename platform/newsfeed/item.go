package newsfeed

import (
	"encoding/json"
	"encoding/xml"
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

// UnmarshalItemJSON returns an item or an error from the given bytes
func UnmarshalItemJSON(data []byte) (Item, error) {
	var r Item
	err := json.Unmarshal(data, &r)
	return r, err
}

// MarshalJSON returns bytes or an error from the given item
func (r *Item) MarshalJSON() ([]byte, error) {
	return xml.Marshal(r)
}

// UnmarshalItemXML returns an item or an error from the given bytes
func UnmarshalItemXML(data []byte) (Item, error) {
	var r Item
	err := xml.Unmarshal(data, &r)
	return r, err
}
