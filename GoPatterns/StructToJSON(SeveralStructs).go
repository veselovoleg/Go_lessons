package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string // `json:"Name"` , if key in json = name
	Last  string
	Id    int
}

func main() {
	person1 := person{
		First: "Oleg",
		Last:  "Veselkin",
		Id:    28091002,
	}

	person2 := person{
		First: "Masha",
		Last:  "Popenkina",
		Id:    26491962,
	}

	personCouple := []person{person1, person2}
	fmt.Printf("Go data: %+v\n", personCouple)
	//$ [{Oleg Veselkin 28091002} {Masha Popenkina 26491962}]

	toJSON, _ := json.Marshal(personCouple)
	fmt.Println("JSON: ", string(toJSON))
	//[{"First":"Oleg","Last":"Veselkin","Id":28091002},{"First":"Masha","Last":"Popenkina","Id":26491962}]
	for count, v := range personCouple {
		fmt.Println(count, v)
		fmt.Println("\t", v.Id)
		/*			0 {Oleg Veselkin 28091002} //(count, v)
		        	 28091002					//("\t", v.Id)
					1 {Masha Popenkina 26491962}//(count, v)
		        	 26491962*/ //("\t", v.Id)
	}

}
