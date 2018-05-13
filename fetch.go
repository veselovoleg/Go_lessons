package main
// Grab URL code
import (
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)

func main() {
	for _,url := range os.Args[1:] {
		resp, err := http.Get(url)/*If we haven't errors, result returns todo
		"resp" structure*/
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
