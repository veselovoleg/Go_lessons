package main
//go run main.go is:open json decoder
import (
    "log"
    "JSONfromWEB/github" //in /go/src/
    "os"
    "fmt"
)

func main(){
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%d тем:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf("#%-5d %8.9s %.55s\n",
            item.Number, item.User.Login, item.Title)
    }
}