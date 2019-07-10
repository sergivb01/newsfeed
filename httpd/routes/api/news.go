package api

import (
	"encoding/json"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/news"
)

func HandleNews(w http.ResponseWriter, r *http.Request) {
	client.CLI.RLock()
	items := client.CLI.Store.Get()
	client.CLI.RUnlock()

	res := struct {
		Found int         `json:"found"`
		News  []news.Item `json:"news"`
	}{
		Found: len(items),
		News:  items,
	}

	b, err := json.Marshal(res)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
