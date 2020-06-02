package main

import (
	"net/http"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	data := []byte("Hello World!") // slice of bytes

	res.Write(data)
}

func main() {

	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":80", handler)

}
