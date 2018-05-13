package main

import (
    "fmt"
    "encoding/json"
)

func main() {

    sliceVar1 := []string{"John", "Andrew", "Robert"}
    sliceVar2, _ := json.Marshal(sliceVar1)
    fmt.Println(string(sliceVar2))
// ["John", "Andrew", "Robert"]

    mapVar1 := map[string]string{"John": "Accepted", "Andrew": "Waiting", "Robert": "Cancelled"}
    mapVar2, _ := json.Marshal(mapVar1)
    fmt.Println(string(mapVar2))
// {"John": "Accepted", "Andrew": "Waiting", "Robert": "Cancelled"}

}

