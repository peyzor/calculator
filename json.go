package main

import (
	"encoding/json"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func responseWithJsonError(w http.ResponseWriter, statusCode int, msg string) {
	type errResponse struct {
		Error string `json:"error"`
	}

	errResp := errResponse{
		Error: msg,
	}
	responseWithJson(w, statusCode, errResp)
}
