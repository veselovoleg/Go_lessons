package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", myController)
	log.Fatal(http.ListenAndServe(":3001", router))
}

func myController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Oleg!"))
}
