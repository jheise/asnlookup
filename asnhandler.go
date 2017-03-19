package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"

	// external
	"github.com/gorilla/mux"
)

func ASNHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["addr"]

	asnrequest := NewASNRequest(address)
	requests <- asnrequest

	response := <-asnrequest.ResponseChan

	asnoutput := NewASNOutput(address, response.Answer)

	output, err := json.Marshal(asnoutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(output))
}
