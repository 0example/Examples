/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//A switch statement takes a specified expression and compares it against a list of multiple cases.
//Once the case matches the expression, it will execute the code block.

package main

import (
	"fmt"
)

func main() {

	grade := "B"

	switch grade {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	default:
		fmt.Println("Default")
	}
}
