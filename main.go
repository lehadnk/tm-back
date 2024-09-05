package main

import (
	"net/http"
)

func getProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {

	http.HandleFunc("GET /my", getProfile)
	http.HandleFunc("GET /profile", getProfile)
	http.ListenAndServe(":8091", nil)
}
