/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Closing the channel signals to consumers of the channel 
//that no more values will be sent on this channel.


package main

import "fmt"

func main() {
	c := make(chan int, 6)

	c <- 3
	c <- 6
	c <- 9

	fmt.Println(<-c)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
//After closing channel you can still read the values but can't send data.
//If you try to send data after closing the channel it will thraw an error message.