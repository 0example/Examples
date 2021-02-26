/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Return values
//	Returning values from a function is another essential concept and
//	can be done by providing a return type after the first pair of parentheses.

package main

import "fmt"

// Returning a value of type int
func add(x int, y int) int {
	return x + y
}

func main() {
	z := add(20, 30)
	fmt.Println("Sum :", z)
}
