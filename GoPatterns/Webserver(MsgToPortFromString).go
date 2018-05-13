package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	/*If we get request to "/*", call handler function*/
	http.ListenAndServe(":9000", nil) //Port adress
}

//Answer to port:9000 request

func handler(w http.ResponseWriter, r *http.Request) { //request handler
	fmt.Fprint(w, "Hello!")
}
