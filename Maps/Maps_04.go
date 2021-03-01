/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Checking if an item exists

package main

import "fmt"

var m = map[string]string{
	"C001": "Amila",
	"C002": "Gayan",
	"C003": "Nishadhi",
	"C004": "Devindhi",
}

func main() {
	// Test if an item exists
	x, _ := m["C003"]
	fmt.Println("C003 : ", x)
}
