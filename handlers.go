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
		Result: params.Number1 * params.Number2,
	}

	responseWithJson(w, http.StatusOK, resp)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Dividend int `json:"dividend"`
		Divisor  int `json:"divisor"`
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJsonError(w, http.StatusBadRequest, "invalid input")
		return
	}

	if params.Divisor == 0 {
		responseWithJsonError(w, http.StatusBadRequest, "divisor can't be zero")
		return
	}

	type response struct {
		Result int `json:"result"`
	}

	resp := response{
		Result: params.Dividend / params.Divisor,
	}

	responseWithJson(w, http.StatusOK, resp)
}

func handleSum(w http.ResponseWriter, r *http.Request) {
	var numbers []int

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&numbers)
	if err != nil {
		responseWithJsonError(w, http.StatusBadRequest, "invalid input")
		return
	}

	type response struct {
		Result int `json:"result"`
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	resp := response{
		Result: sum,
	}

	responseWithJson(w, http.StatusOK, resp)
}
