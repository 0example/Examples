/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Channels can be identified as pipes that connect goroutines and through which you can send and receive values
//with the channel operator <-

package main

import "fmt"

func main() {
	//Channels also must be created before use:
	// Declaring and initializing channels.
	var c1 chan int
	c1 = make(chan int)
	// or
	c2 := make(chan int) //type of the channel is decided by the values they convey. Here it is, int

	fmt.Println(c1, c2) //prints addresses of channels

	go func() { c1 <- 123 }() //send a integer value to a channel c1 using channel operator <-
	chn := <-c1 // receive from channel c1, and assign value to chn
	fmt.Println(chn)

	s:= []int{1,2,3,4,5}
	go sum(s,c2)
	fmt.Println(<-c2) //block the main routine until recieve the value from channel and after receiving value, main routine continue to execute
}

func sum(s []int, c chan int)  {
	var sum int
	for _, num := range s{
		sum+=num
	}
	c <- sum
}

//When the program is executed value is passed from one goroutine to other through the channel successfully.
//By default, sends and receives block until the other side is ready.
//This allows goroutines to synchronize without explicit locks or condition variables.