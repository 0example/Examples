/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Multiple cases can also execute a single code block,
//as shown in the following code block
package main

import "fmt"

func main() {

	grade := "B"

	switch grade {
	case "A":
		fmt.Println("A")
	case "B", "C":
		fmt.Println("B,C")
	case "D", "E":
		fmt.Println("D,E")
	case "F":
		fmt.Println("F")
	default:
		fmt.Println("default")
	}
}
