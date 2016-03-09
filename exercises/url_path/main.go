package main 

import (
	"io"
	"net/http"
	"strings"
)

type x int

func (h x) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s1 := strings.Split(req.URL.Path, "/")
	s2 := strings.Join(s1, " - ")
	res.Header().Set("Content-Types", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+s2+`</h1><br>`)
}

func main() {

	var h x

	mux := http.NewServeMux()
	mux.Handle("/", h)

	http.ListenAndServe(":8080", mux)
}