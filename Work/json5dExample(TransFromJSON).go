package main

import (
    "fmt"
	"encoding/json"
	"log"
)

func main() { 
    type Movie struct {
        Title string
        Year int `json:"released"`
		Color bool `json:"color,omitempty"`//omitempty -don't show data, if false
		Actors []string
	}
	
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, 
			Actors: []string {"Mike","Teodor"}},
		{Title: "Writer", Year: 1989, Color: true, 
			Actors: []string {"Onore","Zamin"}},
		{Title: "It", Year: 1921, Color: false, 
			Actors: []string {"Cotta","Trinidad"}},
	}

	/*data, err := json.Marshal(movies)
	 if err != nil {
		 log.Fatalf("Crush marshaling JSON: %s", err)
	 }
	 fmt.Printf("%s\n",data) //Not readable
	 /*[{"Title":"Casablanca","released":1942,"color":false,"Actors":["Mike","Teodor"]},
	 {"Title":"Writer","released":1989,"color":true,"Actors":["Onore","Zamin"]},
	 {"Title":"It","released":1921,"color":false,"Actors":["Cotta","Trinidad"]}]
*/
	data, err := json.MarshalIndent(movies, "", " ") //Break all data to raws
	if err != nil {
		log.Fatalf("Crush marshaling JSON: %s", err)
	}
	fmt.Printf("%s\n",data) 
}