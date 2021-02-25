/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Go supports to contain only "for loop".
//Usually for loop is used, when we know how often we must loop.

//The most basic type using a single condition
package main

import "fmt"

func main() {

	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}
