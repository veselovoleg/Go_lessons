package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonString := `[{"login":"Oleg", "pin":1536},{"login":"Masha", "pin":1489}]`
	couplePerson := []baseS{}
	json.Unmarshal([]byte(jsonString), &couplePerson)
	fmt.Println(couplePerson)
	//$ [{Oleg 1536} {Masha 1489}]

}

type baseS struct {
	Login string `json:"login"`
	PIN   int    `json:"pin"`
}
