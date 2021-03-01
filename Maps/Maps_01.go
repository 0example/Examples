/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//A map is random collection of key/pairs.
//using the value can quickly retrieved from the key that activates the array as an index.

//Declaration
//A map can be declared using multiple techniques.
//		* Declaration map and then adding objects.
//		* Initializing  the map using some default values.
//		* Declaring the map using the make function

package main

import "fmt"

func main() {

	//Empty map
	var x = map[string]int{}
	fmt.Println(x)

	// Initializing a map
	var y = map[string]int{"Rob": 25, "James": 20}
	fmt.Println(y)

	//Using make function
	var z = make(map[string]int)
	z["Gayan"] = 30
	z["Gayani"] = 25
	fmt.Println(z)
}
