package main

import (
    "fmt"
    "encoding/json"
)

func main() { 

	type animal struct {
		Name string
		Age int
		Food []string
	}

	cat := &animal {
		Name: "Troll",
		Age: 2,
		Food: []string{ "Whiskey", "Kittenkit" }}
	

	cat2, _ := json.Marshal(cat)
	fmt.Println(string(cat2))
    //{"Name":"Troll","Age":2,"Food":["Whiskey","Kittenkit"]}

}