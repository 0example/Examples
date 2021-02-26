/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//By default channels can accept a single send(<-) if there is a receive. They are unbuffered.
//Buffered channels can store multiple values without a corresponding receiver for those values.

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
