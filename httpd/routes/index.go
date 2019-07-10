package routes

import (
	"html/template"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/news"
)

var Templates *template.Template

type indexPage struct {
	Items []news.Item
	Len   int
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	client.CLI.RLock()
	items := client.CLI.Store.Get()
	client.CLI.RUnlock()

	vars := indexPage{
		Items: items,
		Len:   len(items),
	}

	if err := Templates.ExecuteTemplate(w, "index", vars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
