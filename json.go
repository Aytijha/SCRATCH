package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// http.ResponseWriter - the same http response writer that http handlers in Go use (exposed in the library "net/http")
// code int - the status code it will respond with
// payload interface - something we can arrange to a JSON structure
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500) //Internal Service Error
		return
	}
	w.Header().Add("Content-Type", "application/json") //this will say that the response is of this type (application or json)
	w.WriteHeader(200)                                 //everything went well
	w.Write(dat)
}
