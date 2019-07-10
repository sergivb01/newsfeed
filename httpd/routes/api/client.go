package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/sergivb01/newsfeed/httpd/client"
)

func HandleClient(w http.ResponseWriter, r *http.Request) {
	res := struct {
		LoadedSources int       `json:"loadedSources"`
		FetchInterval string    `json:"fetchInterval"`
		LastFetched   time.Time `json:"lastFetched"`
		NextUpdate    time.Time `json:"nextUpdate"`
	}{
		LoadedSources: len(client.CLI.Sources),
		FetchInterval: shortDur(client.CLI.FetchInterval),
		LastFetched:   client.CLI.LastFetched,
		NextUpdate:    client.CLI.NextUpdate,
	}

	b, err := json.Marshal(res)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
