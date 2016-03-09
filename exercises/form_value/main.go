package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, req.FormValue("n"))
	})
	http.ListenAndServe(":8080", nil)
}