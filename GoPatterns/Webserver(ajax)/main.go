package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("home.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {

	err := tpl.Execute(w, "home.html")
	if err != nil {
		log.Fatalln(err)
	}

}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ajax_post_data := r.FormValue("ajax_post_data")
		fmt.Println("При помощи  ajax post получена строка:  ", ajax_post_data)
		w.Write([]byte("<h3>После отправки запроса на сервер<h3>"))
	}
}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/receive", receiveAjax)

	http.ListenAndServe(":8080", mux)
}
