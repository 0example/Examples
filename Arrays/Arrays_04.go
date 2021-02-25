/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Initializing an Array with ellipses
//Automatically detects array size based on the element
//in the array expression using elliptical syntax.

package main

import "fmt"

func main() {

	x := [...]int{10, 20, 30, 40, 50}

	fmt.Println(len(x))
	fmt.Println(x)
}
