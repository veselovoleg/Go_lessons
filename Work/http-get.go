package main

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(body)
	if err != nil {
		log.Fatal(err)
	}
}