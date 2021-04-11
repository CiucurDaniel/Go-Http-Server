package main

import (
	"fmt"
	"net/http"
)

// the bellow functions are Handlers, they take two parameters:
// the http.ResponseWriter w
// and
// a http.Request

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello from GO!\n")
}

// this handler is more special it reads all the HTTP request headers and
// it lists them on the response body

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello from GO index!\n")
}

// register the handlers using the HandleFunc function
// give it the route then the hanlder function for that route
//
// call ListenAndServe with the port and a handler
func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}
