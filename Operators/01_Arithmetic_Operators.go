/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//This operators are used for actions like adding or subtracting numbers.
//
//'+' (Addition) - Adds to operands
//'-' (Subtraction) - Builds the sum of two numbers by subtracting them
//'*' (Multiplication) - Multiplies two values with each other
//'/' (Division) - Divides two values
//'%' (Modulus) - Gets the remainder after an integer division
//'++' (Increment) - It increases the integer value by one
//'--' (Decrement) - It decreases the integer value by one

package main

import "fmt"

func main() {
	var x int = 20
	var y int = 8
	var z int

	z = x + y
	fmt.Printf("x + y = %d\n", z)

	z = x - y
	fmt.Printf("x - y = %d\n", z)

	z = x * y
	fmt.Printf("x * y = %d\n", z)

	z = x / y
	fmt.Printf("x / y = %d\n", z)

	z = x % y
	fmt.Printf("x %% y = %d\n", z)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)

}
