package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, status int, message string, data any) error {
	response := map[string]any{}

	response["status"] = status
	response["data"] = data
	response["message"] = message

	return WriteJsonResponse(w, status, response)
}

func WriteJsonErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{}

	response["status"] = status
	response["error"] = err.Error()
	response["message"] = message

	return WriteJsonResponse(w, status, response)
}

func ReadJsonRequest(r *http.Request, result any) error {
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields() // Prevents unknown fields in the JSON request
	return jsonDecoder.Decode(result)
}
