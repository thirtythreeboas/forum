package pkg

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data := map[string]interface{}{
		"message": msg,
	}

	json.NewEncoder(w).Encode(data)
}

func WriteResponse[T any](w http.ResponseWriter, res T, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		WriteError(w, status, msg)
	}
}
