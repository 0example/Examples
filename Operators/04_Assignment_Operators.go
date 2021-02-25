/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//'='  (Simple Assignment) -  Assigns the value to the variable
//'+=' (Add Assignment) -  Adds the value to the current value of the variable
//'-=' (Subtract Assignment) -  Subtracts the value to the current value of the variable
//'*=' (Multiply Assignment) -  Multiplies the current value of the variable with the new value
//'/=' (Division Assignment) - Divides the current value of the variable by the new value
//'%=' (Modulus Assignment) - Takes modulus of the two values and assigns the result to left operand

package main

import "fmt"

func main() {

	var x, y = 10, 20
	x = y
	fmt.Println(x)

	x = 10
	x += y
	fmt.Println(x)

	x = 10
	x -= y
	fmt.Println(x)

	x = 10
	x *= y
	fmt.Println(x)

	x = 20
	x /= y
	fmt.Println(x)

	x = 20
	y = 3
	x = x % y
	fmt.Println(x)
}
