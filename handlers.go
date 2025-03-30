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
	defer r.Body.Close()

	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJsonError(w, http.StatusBadRequest, "invalid input")
		return
	}

	type response struct {
		Result int `json:"result"`
	}

	resp := response{
		Result: params.Number1 + params.Number2,
	}

	responseWithJson(w, http.StatusOK, resp)
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Number1 int `json:"number1"`
		Number2 int `json:"number2"`
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJsonError(w, http.StatusBadRequest, "invalid input")
		return
	}

	type response struct {
		Result int `json:"result"`
	}

	resp := response{
		Result: params.Number1 - params.Number2,
	}

	responseWithJson(w, http.StatusOK, resp)
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {

}

func handleDivide(w http.ResponseWriter, r *http.Request) {

}

func handleSum(w http.ResponseWriter, r *http.Request) {

}
