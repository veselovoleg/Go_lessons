package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/api/users.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func indexPage(w http.ResponseWriter, r *http.Request) {

	err := tpl.Execute(w, "index.html")
	if err != nil {
		log.Fatalln(err)
	}

}
