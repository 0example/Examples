/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Passing addresses to function
//Pass the address of a variable to a function and modify it with the reference in the function.

package main

import "fmt"

// Passing addresses to a function
func addValues(x *int, y *int) {
	*x = *x + 5
	*y = *y + 10
	return
}

func main() {
	var number1 = 20
	var number2 = 25
	fmt.Println("Before:", number1, number2)

	addValues(&number1, &number2)

	fmt.Println("After:", number1, number2)
}
