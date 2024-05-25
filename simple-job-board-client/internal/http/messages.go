package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ResponseError(w http.ResponseWriter, status int, err error) {
	ResponseJSON(w, status, map[string]string{"error": err.Error()})
}

func ReadJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
