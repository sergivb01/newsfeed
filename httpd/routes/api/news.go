package api

import (
	"encoding/json"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
)

func HandleNews(w http.ResponseWriter, r *http.Request) {
	client.CLI.RLock()
	items := client.CLI.Store.Get()
	client.CLI.RUnlock()

	b, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
