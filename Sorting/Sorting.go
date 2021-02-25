/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Package sort provides primitives for sorting slices and user-defined collections.
//Here are examples for sorting slices.


package main

import (
	"fmt"
	"sort"
)

func main() {
	//Sorting ints
	integers := []int{9, 3, 12, -6}
	sort.Ints(integers)
	fmt.Println("Integers sorted : ",integers)


	//Sorting floats
	floats := []float64{5.8, -1.9, 0.3, 2.7}
	sort.Float64s(floats)
	fmt.Println("Floated sorted : ",floats)

	//Sorting strings
	strings := []string{"mango","orange","apple"}
	sort.Strings(strings)
	fmt.Println("Strings Sorted : ",strings)

	//Checking if slice of integers is sorted
	fmt.Println("Is integers sorted : ",sort.IntsAreSorted(integers))

	//We can also sort a slice of integers in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(integers)))
	fmt.Println(integers)
}