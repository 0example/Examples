/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Iterating an array
//Looping through the array element is done by all sorts of loops.

package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}


	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Getting index and value
	for index, element := range x {
		fmt.Println(index, "-", element)
	}

	// Only getting value
	for _, val := range x {
		fmt.Println(val)
	}

	// Range and counter
	i := 0
	for range x {
		fmt.Println(x[i])
		i++
	}
}