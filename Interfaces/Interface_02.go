/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Implementing the Interface
//	When implementing an interface in your struct you will not need to
//		explicitly specify the interface that is implemented.
//	Go will automatically determine if the struct implements the interface 	by checking
//		if it implements all the method signatures defined in the interface

package main

import (
	"fmt"
)

// Defining a interface
type User interface {
	PrintName()
}

// Create type for interface
type Usr struct {
	id   string
	name string
}

type Car struct {
	name string
}

// Implement interface function
func (usr Usr) PrintName() {
	fmt.Println("User Id:\t", usr.id)
	fmt.Println("User Name:\t", usr.name)

}

// Implement interface function
func (car Car) PrintName() {
	fmt.Println("Car Name:\t", car.name)
}

func main() {
	var x Usr
	x.id = "C001"
	x.name = "Rob"
	x.PrintName()

	var y Car
	y.name = "BMW"
	y.PrintName()
}
