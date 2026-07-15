package handler

import (
	"encoding/json"
	"net/http"
)

type envelope struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func writeJSON(w http.ResponseWriter, code int, body envelope) {
	body.Code = code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(body)
}

func writeSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	writeJSON(w, code, envelope{Status: "success", Message: message, Data: data})
}

func writeError(w http.ResponseWriter, code int, message string) {
	writeJSON(w, code, envelope{Status: "error", Message: message})
}
