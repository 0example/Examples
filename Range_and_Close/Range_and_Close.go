/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//We can use range syntax on channels to iterate over received values.
//You should close the buffered channel before iterate over values.


package main

import "fmt"

func main() {
	c := make(chan int, 6)

	c <- 3
	c <- 6
	c <- 9
	close(c) //We can also close non-empty channels when needed.

	for element := range c{
		fmt.Println(element)
	}
}

//Here, range iterates over received elements from the channel 
//and stops iterating because the channel is closed after receiving 2 values.

