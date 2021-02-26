/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Variadic functions
//	A variadic function can take an arbitrary number of arguments.

package main

import (
	"fmt"
)

func main() {
	multipleStrings("Rob", "James", "Arthur")
}

// Passing multiple attributes using a variadic function
func multipleStrings(m ...string) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}
