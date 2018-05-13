package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) //"map" contains couples "key-value" (str/int)
	//"make" create new clear mapping
	input := bufio.NewScanner(os.Stdin) //Read & Split to the rows our text
	//return TRUE, if row was read
	//return FALSE, if row wasn't read (input data was finished)

	for input.Scan() {
		counts[input.Text()]++ //If we haven't see same key(in previous keys),int values = 0
		/* "[input.Text()]++" is equivalent to
		 " line := input.Text()
		 	counts[line] = counts[line] + 1 "	*/
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) //"Printf" - formatted output
			//  %d - Dec number, %s - String, \t - tabulatin, \n - new row
		}
	}
}
// go run findDuplicateStrings1.go < 1.txt
