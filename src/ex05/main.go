package main

import (
	"github.com/qqqjong/go_http/src/ex05/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}