package api

import (
	"encoding/json"
	"net/http"

	"github.com/sergivb01/newsfeed/httpd/client"
)

func HandleClient(w http.ResponseWriter, r *http.Request) {
	client := client.CLI

	b, err := json.Marshal(client)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
