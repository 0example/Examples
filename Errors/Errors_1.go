/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Go language does not have an exception handling mechanism like try-catch in Java.
//So we cannot throw an exception in Go.
//Go has a different approach for handling errors which is called defer-panic-and-recover mechanism.

//For most go functions and methods, errors are returned last as a normal return value
//and have type error, a built-in interface.
//The error built-in interface type is the conventional interface for representing an error condition,
//with the nil value representing no error.

//An idiomatic way to handle an error is to check the last return value of a function call
//and check for the nil condition.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//ioutil.ReadFile reads the file and and returns the content as a byte slice
	//A successful call returns err == nil
	bs, err := ioutil.ReadFile("00example.txt")

	//If there was an issue reading the file like, filename or filepath is wrong
	//or doesn't have enough permission to read the file, an error value will return
	//after calling the function.
	if err != nil {
		//If the error != nil, the log.Fatal function terminates the program
		//by calling os.Exit() after printing the error to the console
		log.Fatal(err)
	}
	//If err == nil the contents will be printed to the console.
	fmt.Println(string(bs))
}

//OUTPUT : 2021/02/23 19:03:44 open 00example.txt: The system cannot find the file specified.
