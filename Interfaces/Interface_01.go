/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//A set of methods signatures is defined by an interface in golang.

//Defining an Interface
//Using Interface keyword you can define an interface after the type name as given below.

package main

// Defining a interface
type User interface {
	PrintName(name string)
}

type Vehicle interface {
	Alert() string
}

func main() {

}

//The User interface defines a single method signature called PrintName().
//Any type that implements the PrintName() function will therefore now satisfy the User interface
