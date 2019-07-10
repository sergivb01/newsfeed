package api

import (
	"encoding/json"
	"net/http"
)

type errorHandle struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	b, err := json.Marshal(&errorHandle{
		Message: err.Error(),
		Data:    "🖖🏻",
	})

	if err != nil {
		// we have really fucked now lol
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
	w.WriteHeader(statusCode)
}
