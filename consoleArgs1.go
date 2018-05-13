package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { //len - quantity of elements
		s += sep + os.Args[i]

		fmt.Print(i)
		fmt.Print(" ")
		fmt.Print(s)
		fmt.Println(" ")

		s = ""

	}

	sep = " "
}
