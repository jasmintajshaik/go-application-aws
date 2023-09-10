package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloRequest)
	http.ListenAndServe(":8080", nil)

}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
	return
}
func helloRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There!")
	return
}
