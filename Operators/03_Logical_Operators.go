/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//'&&' (Logical AND) - Returns true if all conditions are true
//'||' (Logical OR) -  Returns true if at least one condition is true
//'!'  (Logical NOT) -  Inverts the result of the condition

package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20
	var z int = 30

	fmt.Println(x < z && x > y)
	fmt.Println(x < z || x > y)

}
