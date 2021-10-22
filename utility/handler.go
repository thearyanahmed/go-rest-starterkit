package utility

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
)

// Headers set header to request
func Headers(r http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(r)
}

// Response will return json response of http
// This func handle both error a well as success
func Response(w http.ResponseWriter, payload interface{}, httpCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(httpCode)

	json.NewEncoder(w).Encode(payload)
}

func ReadBody(r *http.Request, data interface{}) (interface{}, error) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	return data, err
}

func SuccessPayload(data interface{}, message string, args ...int) map[string]interface{} {
	result := make(map[string]interface{})

	result["data"] = data
	result["message"] = message

	return result
}
