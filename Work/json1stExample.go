package main

import (
    "fmt"
    "encoding/json"
)

func main() {

type User struct {
    FirstName string
    LastName string
    Books [] string
}
userVar1 := &User{
    FirstName: "John",
    LastName: "Smith",
    Books: []string{ "The Art of Programming", "Golang for Dummies" }}
userVar2, _ := json.Marshal(userVar1)

fmt.Println(string(userVar2))
}
/*{"FirstName":"John","LastName":"Smith","Books":["The Art of Programming",
    "Golang for Dummies"]}*/
