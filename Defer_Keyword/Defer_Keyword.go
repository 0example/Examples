/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Defer is a special keyword in go used to postpone the execution of a function or a statement until the current 
//function has completed.
//It is mainly used for cleaning purposes.


package main

import (
	"fmt"
)

func main() {
	defer print1("Hello")
	print2("World")
}

func  print1(str string)  {
	fmt.Println(str)
}

func print2(str string)  {
	fmt.Println(str)
}