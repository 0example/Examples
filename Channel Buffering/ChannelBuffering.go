/*	Copyright 2021 www.0example.com, powered by Wixis360  */

package main

import "fmt"

func main() {
	var c chan string = make(chan string, 3) // creates a buffered channel with the capacity of 3

	c <- "Go"
	c <- "Lang"
	c <- "Oexample" //can store multiple values like this in the buffered channel.

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
