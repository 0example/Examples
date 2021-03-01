/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Using select with default case we can receive a value from a channel if a value is available.
//If there is no value the default case will be executed
package main

import "fmt"

func main() {
	chn := make(chan string)

	select {
	case status := <-chn:
		fmt.Println("Status received :", status)
	default:
		fmt.Println("no status received")
	}
}
//OUTPUT: no status received