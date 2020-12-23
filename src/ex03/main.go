package main

import (
	"net/http"

	"github.com/qqqjong/go_http/src/ex03/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}