package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
 	counts := make(map[string]int) /*Returns key/value. If we look at new string
	(not repitted), value = 0*/
	files := os.Args[1:] // Open files from args.

	if len (files) == 0 { //quantity of elements, get from command line
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) /*os.Open returns two values. First - opened
			file (*os.File). File will be readded by Scanner
			Second - value of the embedded type "error"*/
				if err != nil { //if err == nil, file was opened correctly
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err)//Error message
					continue //to the next file
				}

			countLines (f, counts)
			f.Close() //The file was read. File closed, all reourses released
		}
	}
	for line, n := range counts {
			if n > 1 { //If we found two or more duplicate strings, print it
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f) //Scanning file
	for input.Scan() {
		counts[input.Text()]++ //Add rows from files
	}

}
//cat 1.txt 2.txt | go run findDuplicateStrings2.go
