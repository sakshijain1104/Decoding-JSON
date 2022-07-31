package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	//some keys used in JSON
	Name     string
	Jobtitle string
	Phone    []string
	Email    string
}

func main() {
	DecodeJson()
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"name": "Mark Taylor",
		"jobtitle": "Software Developer",
		"phone": {
			"home":"123-466-799",
			"office":"564-987-654"
		},
		"email": "markt@gmail.com"
	}
	`)

	//add data to struct format
	var data_info data

	//to check whether JSON is valid or not
	checkValidity := json.Valid(jsonData)

	if checkValidity {
		fmt.Println("Json valid")
		json.Unmarshal(jsonData, &data_info)
		fmt.Printf("\n Way 1 - struct format\n")
		fmt.Printf("%#v\n", data_info)
	} else {
		fmt.Println("Json NOT valid")
	}

	//add data to key value pair
	//using interface as the data coming from json might be array or any other type rather than just string
	fmt.Printf("\n Way 2 - map format\n")
	var mydata map[string]interface{}
	json.Unmarshal(jsonData, &mydata)
	fmt.Printf("%#v\n", mydata)

	//order not fixed in key-value pairs
	fmt.Printf("\n Way 3 - keys and values\n")
	for key, value := range mydata {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, value, value)
	}

}
