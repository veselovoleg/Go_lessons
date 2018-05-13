//http://kari-orders-hub-master.dc.wildberries.ru:6660/api/v1/orders?rv=0
package main

import (
	"database/sql"
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
	fmt.Printf("%s\n", getStringHTML)

	//Unmarhsal HTML body, receive string
	json.Unmarshal(getStringHTML, &showStruct)
	fmt.Println(showStruct.Content)
	log.Printf("%d", showStruct.Rv)

	//MArshal my data, print in console like a   JSON structure
	data, err := json.MarshalIndent(showStruct, "", " ") //Break all data to raws
	if err != nil {
		log.Fatalf("Crush marshaling JSON: %s", err)
	}
	fmt.Printf("%s\n", data)

	/*
		//Write file
		file, err := os.Create("SavedStructure.bin")
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.Write(getStringHTML)*/

	//Connect to DB
	connStr := "user=test_user password=p dbname=testtask sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error of connection with DataBase")
	}
	fmt.Println("DB connected")
	defer db.Close()

	/*
			rows, err := db.Query("SELECT * FROM orders")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()


		//Insert our Data to database
		result, err := db.Exec("INSERT INTO orders VALUES (785, 894);")
		if err != nil {
			fmt.Println("Error. Can not to add data to db")
		}

		fmt.Println(result.RowsAffected()) // количество добавленных строк
	*/
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

/*
type baseStruct struct {
	Rv      int `json:"rv"`
	Content `json:"content"`
}

type Content []struct {
	Goods       int       `json:"goods"`
	Status      string    `json:"status"`
	OrderID     int       `json:"order_id"`
	StoreID     int       `json:"store_id"`
	DateCreated time.Time `json:"date_created"`
}

type Goods []struct {
	Gid    int    `json:"gid"`
	Price  int    `json:"price"`
	Status string `json:"status"`
	ChrtID int    `json:"chrt_id"`
}
*/
