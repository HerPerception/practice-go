package main

import (
	"fmt"
	"net/http"
)

func Handler(write http.ResponseWriter, read *http.Request) {
	fmt.Fprint(write, "This is your first class on ASCII Art Web!")
}

func main() {
	http.HandleFunc("/", Handler)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
