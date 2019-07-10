package api

import (
	"encoding/json"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
	"github.com/sergivb01/newsfeed/news"
)

func HandleSources(w http.ResponseWriter, r *http.Request) {
	sources := client.CLI.Sources

	res := struct {
		Found   int           `json:"found"`
		Sources []news.Source `json:"sources"`
	}{
		Found:   len(sources),
		Sources: sources,
	}

	b, err := json.Marshal(res)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
