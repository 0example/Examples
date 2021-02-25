/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//When values are sent to an unbuffered channel the deadlock error occurred.
//We can avoid this issue by using default case
package main

import "fmt"

func main() {
	chn := make(chan string) //Unbuffered channel.

	status := "ON"
	select {
	case chn <- status:
		fmt.Println("Status set :", status)
	default:
		fmt.Println("Status not set ")
	}
}
//OUTPUT : Status not set