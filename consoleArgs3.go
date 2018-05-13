package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[1:], " ")) //want ([]string, string)
	//Fast variant. We can see only values
}
