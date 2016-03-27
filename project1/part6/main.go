package main

import (
	"github.com/nu7hatch/gouuid"
	"crypto/hmac"
	"crypto/sha256"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"io"
	"strings"
)

type User struct {
	Name string
	Age  string
}

func getKey(data string) string{
	h := hmac.New(sha256.New, []byte("keyVal"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func foo () string{

	m := User{
		Name: "Bob",
		Age: "33",
	}

	bs, err := json.Marshal(m)
		if err != nil{
			fmt.Println(err)
		}

		return string(bs)
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

		currentUser := User{
			Name: pName,
			Age: pAge,
		}

		bs, err := json.Marshal(currentUser)
		if err != nil{
			fmt.Println(err)
		}

		jsonB64 := base64.StdEncoding.EncodeToString(bs)

		cookie, err := req.Cookie("session-fino")
		if err != nil {

			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + "|" + pName + "|" + pAge + "|" + jsonB64,
				//Secure: true
				HttpOnly: true,
			}
			cookie.Value += "|" + getKey(cookie.Value) + "|" + cookie.Value
			http.SetCookie(res, cookie)
		}

		
		xs := strings.Split(cookie.Value, "|")
		usrName := xs[1] + "lmao"
		usrCode := xs[3]
		
		if usrCode == getKey(usrName){
			fmt.Println("valid Cookie")
		} else{
			fmt.Println("not a valid Coookie")
		}



		err = tpl.Execute(res, nil)
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}
		
	})

	http.ListenAndServe(":8080", nil)
}
