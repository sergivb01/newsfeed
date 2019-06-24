package store

import (
	"github.com/sergivb01/newsfeed/news"
)

type Storage interface {
	Get() []news.Item
	Set([]news.Item)
}
