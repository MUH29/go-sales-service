package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//https://go101.org/article/value-conversions-assignments-and-comparisons.html
	h := http.HandlerFunc(Echo)

	log.Println("Listening on port 8080")
	//https://pkg.go.dev/net/http@go1.17.5#ListenAndServe
	err := http.ListenAndServe("localhost:8080", h)
	if err != nil {
		//https://www.practical-go-lessons.com/chap-31-logging
		log.Fatal(err)
	}

}

//https://pkg.go.dev/net/http
func Echo(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(w, "You requested for: ", request.Method, request.URL.Path)
}
