package json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, model any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}

	return json.NewDecoder(r.Body).Decode(model)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
