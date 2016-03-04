package main 

import (
	"io"
	"net/http"
)

type x int

func (h x) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Types", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}

func main() {

	var h x

	mux := http.NewServeMux()
	mux.Handle("/", h)

	http.ListenAndServe(":8080", mux)
}