/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//for User defined Data Type

package main

import (
	"encoding/json"
	"fmt"
)
//Declaring a student type struct
type Student struct {
	ID     string
	Name   string
	Email string
	Contact int
}

func main() {
	//Encoding a struct using marshal function

	//Defining a struct instance
	stEnc := Student{
		ID:     "S001",
		Name:   "James Arthur",
		Email: "james.arthur@gmail.com",
		Contact: 023434324,
	}

	//Encoding stEnc into JSON format
	stB, err := json.Marshal(stEnc)
	if err != nil {
		fmt.Println("error:", err)
	}

	//stB is converted into a string as it is encoded into a byte array
	fmt.Println(string(stB))

	//Decoding a struct using unmarshal funtion

	//Declaring a struct instance
	var stDec Student

	//Decoding stB struct from JSON format
	err2 := json.Unmarshal(stB , &stDec)

	if err != nil {
		fmt.Println(err2)
	}
	//Printing from decoded data
	fmt.Println(stDec.ID, stDec.Name, stDec.Email, stDec.Contact)


}