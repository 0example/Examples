/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Assigning function to a variable
//	Anonymous functions allow you to assign a function to variable.

package main

import "fmt"

//anonymous function
var (
	sum = func(x int, y int) int {
		return x + y
	}
)

func main() {
	sum1 := sum(5, 10)
	fmt.Println(sum1)
}
