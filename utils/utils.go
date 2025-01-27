package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Msg string
}

func (r ErrorResponse) Error() string {
	return string(r.Msg)
}

var Validator = validator.New()

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func JsonMessage(msg string) map[string]string {
	return map[string]string{"message": msg}
}

func GetUrlVariable(variableName string, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	variable, ok := vars[variableName]
	if !ok {
		return "", ErrorResponse{Msg: fmt.Sprintf("missing %s", variableName)}
	}

	return variable, nil
}

func GetQueryParam(paramName string, r *http.Request) (string, error) {
	param := r.URL.Query().Get(paramName)
	if param == "" {
		return "", ErrorResponse{Msg: "Invalid name"}
	}
	return param, nil
}
