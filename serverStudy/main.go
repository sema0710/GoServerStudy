package main

import (
	"net/http"
	"test/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}