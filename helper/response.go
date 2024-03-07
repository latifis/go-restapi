package helper

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON digunakan untuk mengirimkan respon JSON ke client
func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ResponseError digunakan untuk mengirimkan respon error JSON ke client
func ResponseError(w http.ResponseWriter, code int, message string) {
	response := map[string]string{"error": message}
	ResponseJSON(w, code, response)
}
