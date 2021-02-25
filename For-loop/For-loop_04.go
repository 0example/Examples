/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Using continue keyword you can iterate the loop

package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
