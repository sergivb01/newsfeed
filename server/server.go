package server

import (
	"sync"

	"github.com/sergivb01/newsfeed/news"
	"github.com/sergivb01/newsfeed/store"
)

var Server Srv

type Srv struct {
	Something string
	Sources   []news.Source

	sync.Mutex
	store.Storage
}
