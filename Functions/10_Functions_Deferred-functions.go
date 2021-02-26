/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Deferred function
//	Defer is unique Go statement that schedules a function call to be run after the function completes.

package main

import "fmt"

func firstFunction() {
	fmt.Println("First Function")
}

func secondFunction() {
	fmt.Println("Second Function")
}

func main() {

	// The Second function called after the first
	defer secondFunction()
	firstFunction()
}
