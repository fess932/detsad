package main

import (
	"net/http"
	"os"
)

type httpHandler struct{}

func (h httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello World!")
	res.Write(data)
}

func main() {
	handler := httpHandler{}

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, handler)
}
