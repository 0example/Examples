/*	Copyright 2021 www.0example.com, powered by Wixis360  */

package main

import (
	"encoding/json"
	"fmt"
	"log"
)
//Declaring a student type struct
type Student struct {
	SID     string
	Name   string
	Email string
	Contact int
}

//JSON field
//	* We can also specify  optional field tags when declaring a structure.
//	* These will help you when formatting an instance of the struct into a specific format and vise versa.
type School struct {
	ID   string      `json:"id"`
	SchoolName string `json:"schoolName"`
	Location string `json:"location"`
}

func main() {
//Encoding a struct using marshal function

	//Defining a struct instance
	stEnc := Student{
		SID:     "S001",
		Name:   "James Arthur",
		Email: "james.arthur@gmail.com",
		Contact: 023434324,
	}

	//Encoding stEnc into JSON format
	stB, err1 := json.Marshal(stEnc)
	checkErr(err1)

	//stB is converted into a string as it is encoded into a byte array
	fmt.Println(string(stB))

//Decoding a struct using unmarshal function

	//Declaring a struct instance
	var stDec Student

	//Decoding stB struct from JSON format
	err2 := json.Unmarshal(stB , &stDec)
	checkErr(err2)

	//Printing decoded data
	fmt.Println(stDec.SID, stDec.Name, stDec.Email, stDec.Contact)

//Encoding a struct using marshal function

	//Defining a struct instance
	stScl := School{
		ID:   "S001",
		SchoolName: "Golang College",
		Location: "Google",
	}

	//Encoding stScl into JSON format
	sclB, err3 := json.Marshal(stScl)
	checkErr(err3)

	fmt.Println(string(sclB))

//Decoding a struct using unmarshal function

	//Declaring a struct instance
	var sclDec School

	//Decoding sclB struct from JSON format
	err4 := json.Unmarshal(sclB , &sclDec)
	checkErr(err4)

	//Printing decoded data
	fmt.Println(sclDec.ID, sclDec.SchoolName, sclDec.Location)

}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}