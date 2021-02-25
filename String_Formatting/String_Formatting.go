/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//printf in go is similar to the C language printf. 
//The format 'verbs' are derived from C's but are simpler.           '
 
//Go provides some verbs to format the strings 
//Here, are some verbs.

package main

import "fmt"

func main() {
	type value struct {
		a,b int

	}

	//General Verbs:
	val := value{a:3,b:6}

	//%v	the value in a default format
	fmt.Printf("%v\n", val)

	//when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v\n", val)

	//%#v	a Go-syntax representation of the value
	fmt.Printf("%#v\n", val)

	//%T	a Go-syntax representation of the type of the value
	fmt.Printf("%T\n", val)

	//%%	a literal percent sign; consumes no value
	fmt.Printf("%%\n")

	//Boolean Verbs
	//%t	the word true or false
	fmt.Printf("%t\n", false)

	//Integer Verbs:
	//	%b	base 2 == represent as binary
	fmt.Printf("%b\n", 8)
	//	%c	the character represented by the corresponding Unicode code point
	fmt.Printf("%c\n", 65)

	//	%d	base 10
	fmt.Printf("%d\n", 26)

	//	%o	base 8
	fmt.Printf("%o\n", 20)

	//	%O	base 8 with 0o prefix
	fmt.Printf("%O\n", 20)

	//	%q	a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%q\n", "\"string\"")

	//	%x	base 16, with lower-case letters for a-f
	fmt.Printf("%x\n", "example")

	//to represent of a pointer, use %p.
	fmt.Printf("%p\n", &val)

	//FloatingPoints
	//Printing a floating-point number
	fmt.Printf("%g", 12.11)

}