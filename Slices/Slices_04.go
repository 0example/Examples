/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Changing item values
//Values can be changed by pointing to a specific index and
//setting a new value.

package main

import "fmt"

func main() {
	// Change item values
	var a = []int{10, 20, 30, 40, 50}
	fmt.Println(a)
	a[0] = 15
	a[1] = 25
	a[2] = 35
	fmt.Println(a)
}
