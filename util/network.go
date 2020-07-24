package util

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ResponseError(w http.ResponseWriter, status int, payload interface{}) {
	ResponseJSON(w, status, map[string]interface{}{
		"message": payload,
	})
}
