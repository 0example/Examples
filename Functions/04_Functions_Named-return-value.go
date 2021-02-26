/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Named return value
//We can define functions with named return types in round brackets after the function parameters.

package main

import "fmt"

func increment(x, y, z, r int) (a, b, c, d int) {

	a = x + 1
	b = y + 1
	c = z + 1
	d = r + 1
	return
}

func main() {

	a, b, c, d := increment(1, 10, 100, 1000)

	fmt.Println(a, b, c, d)
}
