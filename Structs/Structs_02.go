/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Creating a struct

//There are multiple ways to create a structure.
//Here is a list of such common design process.
//	* Declaring a variable with the struct as the datatype without initializing it.
//	* Creating an instance using struct literate.
//	* Creating an instance  using the  new keyword
//	* Using the pointer address operator to create a new instance.

package main

import "fmt"

// Declaring a Struct
type Person struct {
	name string
	age  int
}

func main() {
	//instance of a struct
	var p1 Person
	p1.name = "Rob"
	p1.age = 40
	fmt.Println(p1)

	//using struct literate
	var p2 = Person{name: "James", age: 35}
	fmt.Println(p2)

	//using the new keyword
	var p3 = new(Person)
	p3.name = "Arthur"
	p3.age = 35
	fmt.Println(p3)

	//using the pointer address operator
	var p4 = &Person{name: "Ben", age: 25}
	fmt.Println(p4)
}
