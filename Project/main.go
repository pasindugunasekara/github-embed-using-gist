package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Html/*.gohtml"))
}

func helloWorld(t http.ResponseWriter, x *http.Request) {
	tpl.ExecuteTemplate(t, "index.gohtml", nil)
	// fmt.Fprintf(t, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
