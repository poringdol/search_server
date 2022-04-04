package http

import (
	"encoding/json"
	"net/http"
	"yandex-food/utils"
)

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func NewHTTPError(code int, message string) *utils.HHTPError {
	return &utils.HHTPError{Code: code, Message: message}
}

func SetHTTPError(w http.ResponseWriter, code int, message string) {
	httpErr := NewHTTPError(code, message)
	response := utils.HTTPResponse{Error: httpErr}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
