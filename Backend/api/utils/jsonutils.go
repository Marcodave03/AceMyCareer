package utils

import (
    "net/http"
    "encoding/json"
)

func WriteJson( w http.ResponseWriter, status int, value any) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(value)
}
