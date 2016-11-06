package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const baseUrl = "http://www.amazon.de/gp/product/"

// GetMovie handles the request. Crawls amazon prime page for the movie specified
// by the amazon_id, finds information about the movie and returns it in the response.
func GetMovie(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	aid := params["amazon_id"]

	respCode := http.StatusOK
	resp, err := http.Get(baseUrl + aid)
	if err != nil || !isValid(resp.StatusCode) {
		respCode = http.StatusBadRequest
		writeResponse(nil, respCode, w)
		return
	}
	defer resp.Body.Close()

	m, err := ParseDom(resp.Body)
	if err != nil {
		respCode = http.StatusInternalServerError
	}

	writeResponse(m, respCode, w)
}

func writeResponse(data interface{}, respCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respCode)
	if respCode == http.StatusOK {
		json.NewEncoder(w).Encode(&data)
	}
}

func isValid(statusCode int) bool {
	code := statusCode / 100
	if code == 2 || code == 3 {
		return true
	}

	return false
}
