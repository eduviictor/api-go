package main

import (
	"fmt"
	"net/http"
)

func helloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!\n")
}

func main() {
		http.HandleFunc("/", helloWorld)
		
		http.ListenAndServe(":3000", nil)
}