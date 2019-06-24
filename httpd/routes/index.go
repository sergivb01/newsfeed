package routes

import (
	"html/template"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/news"
)

var Templates *template.Template

type indexPage struct {
	Items   []news.Item
	Sources map[string]news.Source
	Len     int
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	items := client.CLI.Store.Get()
	vars := indexPage{
		Items:   items,
		Sources: client.CLI.Sources,
		Len:     len(items),
	}

	if err := Templates.ExecuteTemplate(w, "index", vars); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Write([]byte("Check out /api/sources, /api/news and /api/client!"))
}
