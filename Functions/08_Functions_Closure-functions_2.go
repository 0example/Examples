/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//This function addValues returns another function, which we define anonymously in the body of addValues.
//The returned function closes over the variable z to form a closure.

package main

import "fmt"

func addValue() func() int {

	return func() int {
		var x int = 5
		var y int = 10
		var z =x+y
		return z
	}
}

func main() {
	sum := addValue()
	fmt.Println(sum())
}
