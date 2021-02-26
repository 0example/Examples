/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Comparing struct instances
//	* "==" operator or the reflect.DeepEqual() method is used to compare
//		 two instances of the same structure with each other.
//	* If all the values are the same then the truth will be returned.

package main

import (
	"fmt"
	"reflect"
)

// Declaring a Struct
type Person struct {
	name string
	age  int
}

func main() {
	// Creating an instance
	var p1 = new(Person)
	p1.name = "Namal"
	p1.age = 25
	fmt.Println(p1)

	var p2 = &Person{name: "Kasun", age: 30}
	fmt.Println(p2)

	// Comparing struct instances
	fmt.Println(p1 == p2)

	// Comparing struct instances using DeepEqual
	fmt.Println(reflect.DeepEqual(p1, p2))
}
