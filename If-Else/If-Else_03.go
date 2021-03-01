/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//If statement can also be chained using an optional else if.. else statement as the following example

package main

import (
	"fmt"
)

func main() {

	marks := 65

	if marks >= 75 {
		fmt.Println("A")
	} else if marks >= 65 {
		fmt.Println("B")
	} else if marks >= 55 {
		fmt.Println("C")
	} else if marks >= 45 {
		fmt.Println("S")
	} else {
		fmt.Println("F")
	}
}
