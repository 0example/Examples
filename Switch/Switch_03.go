/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Switch also allows you to use conditional statements in your cases.
package main

import (
	"fmt"
)

func main() {

	number := 9
	switch {
	case number < 10:
		fmt.Println("Smaller than 10")
	case number == 10:
		fmt.Println("Ten")
	case number > 10:
		fmt.Println("Bigger than 10")
	default:
		fmt.Println("No information about the number")
	}
}
