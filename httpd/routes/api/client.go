package api

import (
	"encoding/json"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
)

func HandleClient(w http.ResponseWriter, r *http.Request) {
	items := client.CLI

	b, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
