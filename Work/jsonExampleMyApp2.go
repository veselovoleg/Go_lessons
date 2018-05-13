package main

import (
    "fmt"
    "encoding/json"
)

func main() {
	
	type user struct  { 
		FirstName string 
		SecondName string 
		Age int 
		ShoppedItems [] string
		}

	user1 := &user {
		FirstName: "Pavel",
		SecondName: "Kurdenko",
		Age: 27,
		ShoppedItems: []string {"Vantus", "Bread", "Sigarets"}}
	

	user1JSON, _ := json.Marshal(user1) 
	fmt.Println(string(user1JSON))

}