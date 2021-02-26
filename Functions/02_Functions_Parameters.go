/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Call By Value
//	When parameters passing, values of parameters are copied to function’s new parameters
//	and the two types of parameters are stored in different memory locations.
//	So any changes made inside functions are not reflected in actual parameters of the caller.

package main

import "fmt"

func hello(x string) {
	fmt.Println(x)
}

func addInt(x int, y int) {
	fmt.Println(x + y)
}

func change(x, y int) {
	x, y = 100, 200
	fmt.Printf("Inside the change func: x = %v AND y = %v\n",x,y)
}
func main() {
	hello("Hello World!")

	var x, y = 10, 20
	addInt(x, y)

	change(x, y)

	fmt.Printf("x = %v AND y = %v",x,y)
}
