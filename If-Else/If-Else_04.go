/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//In golang a short variable can be declared before the conditional expression in an if statement
package main

import (
	"fmt"
)

func main() {
	if a := 15; a == 15 {
		fmt.Println(a)
	}
}
