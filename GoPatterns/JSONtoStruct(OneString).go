package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var dat baseStruct
	jsonString := `{"login":"Oleg", "pin":1536}`
	json.Unmarshal([]byte(jsonString), &dat)
	fmt.Println(dat)

}

type baseStruct struct {
	Login string `json:"login"`
	PIN   int    `json:"pin"`
}
