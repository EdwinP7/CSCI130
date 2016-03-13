package main

import(
	"log"
	"html/template"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}