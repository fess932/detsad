package main

import (
	"net/http"
)

type httpHandler struct{}

func (h httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}

func main() {
	handler := httpHandler{}
	http.ListenAndServe(":80", handler)
}
