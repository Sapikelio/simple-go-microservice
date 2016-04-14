package main

import (
	"./model"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, user.GetUser().String())
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/", getUser)
	http.ListenAndServe(":8080", nil)
}
