/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Infinite loop is shown in this example. We can stop the loop by using the break statement
package main

import (
	"fmt"
)

func main() {

	i := 0
	for {
		fmt.Println("Hello World")
		// Breaks Stops the infinite loop
		if i == 10 {
			break
		}
		i++
	}
}
