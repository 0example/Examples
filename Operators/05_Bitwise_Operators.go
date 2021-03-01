/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//'&'		(AND Operator) - a&b - Returns a one in each bit position for which the corresponding bit of both operands are ones
//'|'		(OR Operato) - a|b - Returns a one in each bit position for which the corresponding bits of either or both operands are ones
//'^'		(XOR Operator) - a^b - Returns a one in each bit position for which the corresponding bits of either but not both operands are ones
//'<<'	(Left Shift Operator) - a<<b - Shift a in binary representation b bits to the left, shifting in zeroes from the right
//'>>'	(Right Shift Operator) - a>>b - Shift a binary representation b bits to the right, discarding bits shifted off.

package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20
	var z int = 0

	z = x & y
	fmt.Println(z)

	z = x | y
	fmt.Println(z)

	z = x ^ y
	fmt.Println(z)

	z = x << 1
	fmt.Println(z)

	z = x >> 1
	fmt.Println(z)
}
