/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Go language has built in support for base64 encoding.
//Encoding means converting information from one form to another, and can be converted to the previous form
//using the same algorithm that encoded it.
//The base64 encoding uses a symbol table of 64 characters to encode the data into a string containing those symbols.
//The 64 characters are A-Z, a-z, 0-9, +, /, =

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	//String to encode
	str := "https://www.0example.com"

	//takes a byte slice and returns a base64 encoding as a string.
	encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("Encoded Data :", encodedStr)

	//Decodes the given base64 encoded string and return a byte slice
	//will return an error if the input string is invalid.
	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic("Invalid input")
	}
	fmt.Println("Decoded Data :", string(decodedStr))
}
