package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
 	counts := make(map[string]int)
	files := os.Args[1:]
	if len (files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) /*os.Open returns two values. First - opened
			file (*os.File). File will be readded by Scanner
			Second - value of the embedded type "error"*/
			if err != nil { //if err == nil, file was opened correctly
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)//Error message
				continue
			}
			countLines (f, counts) //to the next file
			f.Close() //The file was read. File closed, all reourses released
		}
		}
	for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
}

func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
//cat 1.txt 2.txt | go run findDuplicateStrings2.go
