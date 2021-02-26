/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Copying a struct
//	* Here we are copying a structure, when we need a second case,
//		with values without giving the first opportunity.
//	* It feels like changing one situation does not affect the other.

package main

import "fmt"

// Declaring a Struct
type Person struct {
	name string
	age  int
}

func main() {
	// using the pointer address operator
	var p1 = &Person{name: "James", age: 25}
	fmt.Println(p1)

	// Copying struct type using pointer reference
	p2 := p1
	p2.name = "Rob"
	fmt.Println(p2)
}
