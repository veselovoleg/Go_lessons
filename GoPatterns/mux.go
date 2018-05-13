package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/page/{page}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the  page %s\n", page)
		fmt.Println(page)

	})

	http.ListenAndServe(":8080", r)
}
