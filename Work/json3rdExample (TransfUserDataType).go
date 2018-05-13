package main

import (
    "fmt"
    "encoding/json"
)

func main() {

    type User struct {
    FirstName string
    LastName string
    Books []string
}
userVar1 := &User{
    FirstName: "John",
    LastName: "Smith",
	Books: []string{ "The Art of Programming", "Golang for Dummies" }}
	
userVar2, _ := json.Marshal(userVar1)
fmt.Println(string(userVar2))
// {"FirstName":"John","LastName":"Smith","Books":["The Art of Programming","Golang for Dummies"]}

type User2 struct {
    FirstName string `json:"name"` // свойство FirstName будет преобразовано в ключ "name"
    LastName string `json:"lastname"` // свойство LastName будет преобразовано в ключ "lastname"
    Books []string `json:"ordered_books"` // свойство Books будет преобразовано в ключ "ordered_books"
} 

userVar3 := &User2{ 
    FirstName: "John", 
    LastName: "Smith", 
	Books: []string{ "The Art of Programming", "Golang for Dummies" }} 
	
userVar4, _ := json.Marshal(userVar3) 
fmt.Println(string(userVar4)) 
// {"name":"John","lastname":"Smith","ordered_books":["The Art of Programming","Golang for Dummies"]}

}