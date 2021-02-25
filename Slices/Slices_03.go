/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Accessing items
//As arrays, using slices can be indices to access values.

package main

import (
	"fmt"
)

func main() {
	//Access slice items
	var a = []int{10, 20, 30, 40, 50}
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[1:4])
}

