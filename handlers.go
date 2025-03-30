package main

import (
	"encoding/json"
	"net/http"
)

func handleAdd(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Number1 int `json:"number1"`
		Number2 int `json:"number2"`
	}

	decoder := json.NewDecoder(r.Body)

	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input"))
		return
	}

	type response struct {
		Result int `json:"result"`
	}

	resp := response{
		Result: params.Number1 + params.Number2,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {

}

func handleMultiply(w http.ResponseWriter, r *http.Request) {

}

func handleDivide(w http.ResponseWriter, r *http.Request) {

}

func handleSum(w http.ResponseWriter, r *http.Request) {

}
