package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)



func ReadJSON[T any](r io.ReadCloser) (T, error) {
	var v T
	// decode the JSON into v
	err := json.NewDecoder(r).Decode(&v)
	return v, errors.Join(err, r.Close())
}
func WriteJSON(w http.ResponseWriter, v any) {
    w.Header().Set("Content-type", "application/json")
    err := json.NewEncoder(w).Encode(v)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


