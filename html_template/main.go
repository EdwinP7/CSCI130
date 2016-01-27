package main 

import (
	"log"
	"os"
	"text/template"
)

type person struct{
	Name string
	Gender string
}

type check struct{
	person
	Rural bool
}

func main() {
	p1 := check{
		person: person{
			Name: "Sam",
			Gender: "Female",
		},
		Rural: true,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}