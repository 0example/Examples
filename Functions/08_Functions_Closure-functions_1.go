/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Closure functions
//	Closure functions are an anonymous function where you access outside variable.
package main

import "fmt"

func main() {
	x := 5
	y := 10

	// Closure functions are a special case of a anonymous function
	//where you access outside variables
	func() {
		var sum int
		sum = x + y
		fmt.Println(sum)
	}()
}
