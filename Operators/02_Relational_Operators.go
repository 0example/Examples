/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//These operators are used to compare two values with each other

//'==' (Equal to) - Returns true if x is equal to y
//'!=' (Not equal to) - Returns true if x is not equal y
//'<'  (Lesser than) - Returns true if x is lesser than y
//'<=' (Lesser than Equal to) - Returns true if x is lesser than or equal to y
//'>'  (Greater than) - Returns true if x is greater than y
//'>=' (Greater than Equal to) - Returns true if x is greater than or equal to y

package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20

	p := fmt.Println

	p(x == y)
	p(x != y)
	p(x < y)
	p(x <= y)
	p(x > y)
	p(x >= y)
}
