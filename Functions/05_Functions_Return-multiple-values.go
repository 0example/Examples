/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Go language allows to return multiple values from a function.

package main

import "fmt"

// Returning multiple name values
func setValues(x int, y int) (sum int, multiply int) {
	sum = x + y
	multiply = x * y
	return
}

func main() {
	sum, multiply := setValues(5, 10)
	fmt.Println("Sum: ", sum)
	fmt.Println("Multiply", multiply)
}
