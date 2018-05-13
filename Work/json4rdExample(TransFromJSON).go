package main

import (
    "fmt"
    "encoding/json"
)

func main() { 

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
    		panic(err)
		}
fmt.Println(dat) //map[num:6.13 strs:[a b]]


num := dat["num"].(float64) // для того, чтобы получить из свойства num число
fmt.Println(num) //6.13

strs := dat["strs"].([]interface{}) // для того, чтобы получить массив интерфейсов...
str1 := strs[0].(string) // ... и потом получить из него строку #0
fmt.Println(str1) //a

}