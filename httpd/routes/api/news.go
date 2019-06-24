package api

import (
	"github.com/sergivb01/newsfeed/httpd/client"
	"encoding/json"
	"net/http"
)

func HandleNews(w http.ResponseWriter, r *http.Request) {
	items := client.CLI.Store.Get()

	b, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}