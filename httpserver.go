package main

import(
	"net/http"
	"log"
)

func main() {
	http.HandleFunc ("/", myController)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func myController(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello, Oleg!"))
}