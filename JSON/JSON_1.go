/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Go has built-in support for JSON (JavaScript Object Notation) encoding and decoding and
//it also supports custom datatypes.

//We use the  Marshal function to convert go data types into JSON format.
//Syntax: "func Marshal(v interface{}) ([]byte, error)"  

//We use Unmarshal function to decode JSON data type.
//Syntax: func Unmarshal(data []byte, v interface{}) error

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//A byte array is returned after encoding into JSON format

	//Encoding a boolean
	typBool, _ := json.Marshal(true)
	fmt.Println(string(typBool))

	//Encoding an integer
	typInt, _ := json.Marshal(28)
	fmt.Println(string(typInt))

	//Encoding a float
	typFloat, _ := json.Marshal(22.7)
	fmt.Println(string(typFloat))

	//Encoding a string
	typStr, _ := json.Marshal("0example")
	fmt.Println(string(typStr))

	//Encoding a slice
	slcA := []string{"0example", "golang", "google"}
	slcB, _ := json.Marshal(slcA)
	fmt.Println(string(slcB))

	//Encoding a map
	mapx := map[string]string{"manager": "admin", "customer": "user"}
	mapxB, _ := json.Marshal(mapx)
	fmt.Println(string(mapxB))

	//Decoding the byte array generated after encoding mapx
	var dat map[string]string
	if err := json.Unmarshal(mapxB, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}
