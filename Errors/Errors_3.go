/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//We can create a new error with a formatted error message by using fmt.Errorf() function

package main

import (
	"fmt"
	"math"
	"os"
)

func findSquareRoot(num float64) (float64, error) {
	//We can not calculate square roots for negative numbers
	//because they don't have real square roots
	//So a formatted error message will be returned for negative numbers
	if num < 0 {
		return 0, fmt.Errorf("%0.2f is less than zero", num)
	}
	//For a positive number, it's square root value will be returned.
	return math.Sqrt(num), nil
}

func main() {
	num := 4.0

	//Calling findSquareRoot function to get the square root.
	res, err := findSquareRoot(num)

	//If err != nil it prints the error message and exits the program.
	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	//If the square root is successfully calculated it wil be printed.
	fmt.Printf("SquareRoot of %0.2f = %0.2f\n", num, res)

	num = -4.0
	res2, err2 := findSquareRoot(num)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Printf("SquareRoot of %0.2f = %0.2f\n", num, res2)
}
