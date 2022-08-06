package controller

import (
	"io"
	"net/http"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world")
}
