/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//When the boolean expression is false it is followed by else statement

package main

import (
	"fmt"
)

func main() {

	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is lesser than 5")
	}
}
