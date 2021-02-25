/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//The switch statement also allows you to initialize the variable directly after the switch keyword,
//which will only be accessible inside the statement.

package main

import (
	"fmt"
)

func main() {

	switch number := 10; {
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
