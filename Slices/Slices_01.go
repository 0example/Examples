/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Slices are dynamic arrays that can grow and shrink.
//Like the array, the slices can also be indexed and have a length.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create an empty slice
	var a []int
	fmt.Println(reflect.ValueOf(a).Kind())

	// using the make function
	var b = make([]string, 10, 20)
	fmt.Println("length :", len(b))
	fmt.Println("capacity :", cap(b))

	// Initialize the slice with values using a slice literal
	var c = []int{10, 20, 30, 40}
	fmt.Println("length :", len(c))
	fmt.Println("capacity :", cap(c))
	fmt.Println(c)

	//using the new keyword
	var d = new([50]int)[1:5]
	fmt.Println("length :", len(d))
	fmt.Println("capacity :", cap(d))
	fmt.Println(d)
}
