/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Initializing an array
//	We can initialize an array with pre-defined values using an array literal.
//	An array literal has the number of elements it will hold in square brackets,
//	followed by the type of its elements.
//	After that inside curly braces we can keep the initial values separated by commas of each element
package main

import "fmt"

func main() {
	// Initialize Array using Array literals
	x := [5]int{1, 2, 3, 4, 4}
	var y [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(x)
	fmt.Println(y)
}
