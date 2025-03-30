package main

import "net/http"

func handleAdd(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("add numbers"))
}

func handleSubtract(w http.ResponseWriter, r *http.Request) {

}

func handleMultiply(w http.ResponseWriter, r *http.Request) {

}

func handleDivide(w http.ResponseWriter, r *http.Request) {

}

func handleSum(w http.ResponseWriter, r *http.Request) {

}
