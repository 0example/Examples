/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Array is a collection of elements of a single type.
//It can be used to hold multiple values at once.
//We can get the required value by using an index.
//Arrays can't shrink to grow.

//Array Declaration & Assigning values
//	The data type of the elements holding the array is important
//	after following the number of elements to express the array.
//	The index number indicates the values that can be accessed and assigned in parentheses.

package main

import "fmt"

func main() {

	var myArray [10]int

	// Assigning values
	myArray[0] = 10
	myArray[1] = 2

	fmt.Println(myArray[0])
	fmt.Println(myArray[1])
}
