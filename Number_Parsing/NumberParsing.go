/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//In Go language, we can use package strconv for conversion of strings to numbers.
//Here are some examples for string conversion.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	//converts a string in the given base (0, 2 to 36)
	//and bit size (0 to 64) and returns the corresponding value as an int64 integer
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//Atoi is equivalent to ParseInt(s, 10, 0)
	at, _ := strconv.Atoi("1253")
	fmt.Println(at)

	//When an invalid type is given an error will return.
	val, err := strconv.Atoi("golang")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	//ParseInt can identify hex-formatted numbers and convert them
	h, _ := strconv.ParseInt("0x1f3", 0, 64)
	fmt.Println(h)

	//ParseUint is like ParseInt but for unsigned numbers.
	ui, _ := strconv.ParseUint("7777", 0, 64)
	fmt.Println(ui)

	//converts the string to a floating-point number with the precision specified by bitSize
	fp, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(fp)

}
