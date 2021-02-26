/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//By default channels are bidirectional. 
//This means you can both send and receive data using the channel.
//But you can change the channel direction as receive-only or send-only according to the situation.

package main

import "fmt"

func funcReceive(c <-chan string) { // receive only
	fmt.Println(<-c)
}

func funcSend(c chan<- string) { //send only
	c <- "0example.com"
}

func main() {
	cr := make(chan string, 1)
	cs := make(chan string, 1)

	cr <- "hello, this is"
	funcReceive(cr)

	funcSend(cs)
	fmt.Println(<-cs)
}
