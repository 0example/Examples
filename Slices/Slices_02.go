/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Adding items
//	* Add items to a slice using the append() method. and If takes a slice and a value an argument.
//	* If there is enough space in the slice below, the value of the slice is saved.
//		If not a new slice is created to store the value.

package main

import (
	"fmt"
)

func main() {
	// Add items using the append function
	x := make([]int, 1, 10)
	fmt.Println("length :", len(x))
	fmt.Println("capacity :", cap(x))
	fmt.Println(x)
	x = append(x, 20)
	fmt.Println(x)
}