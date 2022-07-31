// This package demonstrates go docs.
//
// This will decode a json file
package main

import (
	"encoding/json"
	"fmt"
)

// Employee is the struct of the JSON
type Employee struct {
	//some keys used in JSON
	Name     string `json:"name"`
	Jobtitle string `json:"jobDescription"`
	Phone    string `json:"homePhone"`
	Email    string `json:"emailId"`
}

// Reading represents an initialisation of Employee struct
var Reading Employee

// main function calls the decodingJson function
func main() {
	DecodingJson()
}

//DecodingJson prints the decoded jsonString into the Employee struct
func DecodingJson() {
	jsonString := `{"name": "Mark Taylor", 
		"jobDescription": "Software Developer", 
		"homePhone":"123-466-799", 
		"emailId": "markt@gmail.com"
	}`
	json.Unmarshal([]byte(jsonString), &Reading)
	fmt.Printf("%+v\n", Reading) //%+v prints struct with field names and values
}
