/*	Copyright 2021 www.0example.com, powered by Wixis360  */


package main

import "fmt"

func main() {

	//Declaring variables
	var i int
	var s string

	//Initializing variables
	i = 20
	s = "String value"

	fmt.Println(i)
	fmt.Println(s)

	//declaring and Initializing variables in single line
	var x int = 35
	fmt.Println(x)

	//Short variable declaration
	y := 50
	z := "String value"

	fmt.Println(y)
	fmt.Println(z)

	//Declaring multiple variables
	var q, r int = 10, 20
	fmt.Println(q, r)

	//Variable Declaration Block
	var (
		name = "James"
		age  = 25
	)
	fmt.Println(name)
	fmt.Println(age)

	//Go will guess the type of initialized variable
	var t = true
	fmt.Println(t)

	//If a variable is not initialized it gives the default value
	var e int
	fmt.Println(e)
	var k bool
	fmt.Println(k)

}