/*	Copyright 2021 www.0example.com, powered by Wixis360  */

package main

import "fmt"

func main() {
	var x int = 20
	var y int = 8
	var z int

	z = x+y
	fmt.Printf("x + y = %d\n",z)

	z = x-y
	fmt.Printf("x - y = %d\n",z)

	z = x*y
	fmt.Printf("x * y = %d\n",z)

	z = x/y
	fmt.Printf("x / y = %d\n",z)

	z = x%y
	fmt.Printf("x %% y = %d\n",z)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n",y)

}