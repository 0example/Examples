/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//A pointer is a variable which holds the memory address of another variable
//The type *p is a pointer to a p value.
//The &p syntax gives the memory address of p
//The zero value of a pointer is nil.

package main

import "fmt"

func zeroValue(a int) {
	a = 0
}

//Assigning a value to a pointer changes the value
//at the pointer referenced address.
func zeroPointer(p *int) {
	*p = 0 //set x through pointer p
}

func main() {
	x := 1
	fmt.Println("value of x at the beginning:", x)

	zeroValue(x)
	fmt.Println("after calling func zeroValue, x : ", x)

	//The & operator generates a pointer to its operand.
	zeroPointer(&x)
	fmt.Println("After calling func zeroPointer, x :", x)

	//Prints the memory address of x.
	fmt.Println("Pointer address : ", &x)
}
