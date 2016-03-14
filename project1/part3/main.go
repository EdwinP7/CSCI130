package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/nu7hatch/gouuid"
	//"fmt"
)

type Person struct {
	Name string
	Age  string
}

func main() {

	// parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		
		pName := req.FormValue("name")
		pAge := req.FormValue("age")
		

		cookie, err := req.Cookie("session-fino")
		if err != nil{
			
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-fino",
				Value: id.String() + pName + pAge,
				//Secure: true
				HttpOnly: true,
			}

			http.SetCookie(res, cookie)
		}
			err = tpl.Execute(res, nil)
			if err != nil {
				http.Error(res, err.Error(), 500)
				log.Fatalln(err)
			}
	})

	http.ListenAndServe(":8080", nil)
}