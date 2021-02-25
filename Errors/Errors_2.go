/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//errors.New function can create a new Error and it returns an error
//that formats as the given text.

package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(x int, y int) (int, error) {
	//If the denominator is zero, it generates a new error.
	if y == 0 {
		return 0, errors.New("can not be divided by 0")
	} else {
		return x / y, nil
	}
}

func main() {
	var x, y = 8, 4
	if res, err := divide(x, y); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d / %d = %d\n", x, y, res)
	}

	x, y = 4, 0
	if res, err := divide(x, y); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d / %d = %d\n", x, y, res)
	}
}
