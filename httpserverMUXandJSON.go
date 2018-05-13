package main

import (
	"encoding/json"
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
	req := Req{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error":"` + err.Error() + `"}`))
		return
	}
}

type Req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
