/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Copying and appending items
//Slices can be copied using the build-in copy() function,
//Using the append() function and the spread operator we can add all the items
//that have a slice to one another.

package main

import (
	"fmt"
)

func main() {
//Copy slice into another slice
var x = []int{1, 2, 3, 4, 5}
var y = []int{6, 7, 8, 9, 10}
copy(x, y)
fmt.Println(x)

// Append a slice
var p = []int{1, 2, 3, 4, 5}
var q = []int{6, 7, 8, 9, 10}

p = append(p, q...)
fmt.Println(p)
}



