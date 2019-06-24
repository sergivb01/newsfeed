package routes

import (
	"encoding/json"
	"net/http"

	"github.com/sergivb01/newsfeed/server"
)

func HandleNews(w http.ResponseWriter, r *http.Request) {
	items := server.Server.Get()

	b, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
