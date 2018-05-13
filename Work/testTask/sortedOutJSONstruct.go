package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	var showStruct baseStruct
	getStringHTML := getPage() //Received page
	//getStringHTML = nil        //I dont want to see all in console
	//fmt.Printf("%s\n", getStringHTML)

	//Unmarhsal HTML body, receive string
	json.Unmarshal(getStringHTML, &showStruct)
	/*fmt.Println(showStruct.Content)
	log.Printf("%d", showStruct.Rv)*/

	//MArshal my data, print in console like a   JSON structure
	data, err := json.MarshalIndent(showStruct, "", " ") //Break all data to raws
	if err != nil {
		log.Fatalf("Crush marshaling JSON: %s", err)
	}
	fmt.Printf("%s\n", data)

	gid string
		value int
		x map[gid]value

	for i:=0, i < len(data); i++ {
		
	}

}

//URL ..
const URL = "http://kari-orders-hub-master.dc.wildberries.ru:6660/api/v1/orders?rv=0"

func getPage() []byte { //if i have string, i will do Unmashaling

	//fmt.Printf("WebPage: %s ...\n", URL) //Show link. what we used
	resp, err := http.Get(URL)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// close
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
}

//https://mholt.github.io/json-to-go/ - converter
type baseStruct struct {
	Rv      int `json:"rv"`
	Content []struct {
		Goods []struct {
			Gid    int    `json:"gid"`
			Price  int    `json:"price"`
			Status string `json:"status"`
			ChrtID int    `json:"chrt_id"`
		} `json:"goods"`
		Status      string    `json:"status"`
		OrderID     int       `json:"order_id"`
		StoreID     int       `json:"store_id"`
		DateCreated time.Time `json:"date_created"`
	} `json:"content"`
}
