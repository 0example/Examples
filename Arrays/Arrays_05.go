/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Copying an array
//Array copying is done in two ways. We can copy all its values or copy the reference of the array.
//	* copy values - copied arrays have the same values but are independent of the original array.
//	* copy reference - on the other hand, changes in the original array are reflected in the copied array.

package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}

	// Copy array values
	y := x

	// Copy reference
	z := &x

	fmt.Println(x)
	fmt.Println(y)

	x[0] = 1

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)
}
