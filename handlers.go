package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API endpoint :D")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata Metadada
	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload: %v", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	response, err := user.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
