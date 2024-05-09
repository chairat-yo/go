package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	data := []string{"a", "b", "c"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	err = tpl.Execute(nf, data)
	if err != nil {
		log.Fatalln(err)
	}

	person := struct {
		fname string
		lname string
	}{
		"Chairat",
		"Kong",
	}

	fmt.Println(person)
}
