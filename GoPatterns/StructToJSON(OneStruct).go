package main

import (
	"encoding/json"
)

func main() {
	type User struct {
		Login string
		PIN   int
	}

	user1 := &User{
		Login: "Foska",
		PIN:   1578,
	}

	user1m, _ := json.Marshal(user1)
	println(string(user1m))

}
