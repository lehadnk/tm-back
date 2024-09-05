package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func getProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	httpMux := reflect.ValueOf(http.DefaultServeMux).Elem()
	finList := httpMux.FieldByIndex([]int{1})
	fmt.Println(finList)

	http.HandleFunc("GET /my", getProfile)
	http.HandleFunc("GET /profile", getProfile)
	http.ListenAndServe(":8091", nil)

}
