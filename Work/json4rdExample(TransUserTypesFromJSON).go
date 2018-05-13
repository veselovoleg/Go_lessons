package main

import (
    "fmt"
    "encoding/json"
)

func main() { 

	type user struct {
		FirstName string
		LastName string
		Books []string
	}

    //user = user1{}
	userJson := "{\"FirstName\":\"John\",\"LastName\":\"Smith\",\"Books\":[\"The Art of Programming\",\"Golang for Dummies\"]}"
	bytes := []byte(userJson)
	json.Unmarshal(bytes, &user)
	fmt.Println(user.FirstName, user.LastName, user.Books)
// John Smith [The Art of Programming Golang for Dummies]

	user2 := User2{}
	userJson2 := "{\"name\":\"John\",\"lastname\":\"Smith\",\"ordered_books\":[\"The Art of Programming\",\"Golang for Dummies\"]}"
	bytes2 := []byte(userJson2)
	json.Unmarshal(bytes2, &user2)
	fmt.Println(user2.FirstName, user2.LastName, user2.Books)
// John Smith [The Art of Programming Golang for Dummies]

}