package utils

import (
	"encoding/json"
	"net/http"
)

func EnableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin","*")
}

func WriteJson(w http.ResponseWriter, v any, status int) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)

    return json.NewEncoder(w).Encode(v)
}
