package main

import (
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	for _, arg :=range os.Args[1:] { //_ - instead of variable
		/*If we started from os.Args[0], 1st argument = path of compiled file
		C:\Users\vesel\AppData\Local\Temp\go-build013435514\b001\exe\consoleArgs2.exe */
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
