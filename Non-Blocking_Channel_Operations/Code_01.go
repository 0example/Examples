/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//By default, reading and writing operations on channels are blocking.
//By using default case with select , we can acquire non-blocking send, receive operations
//and unblocked multi way selects.

//Pushing a value to a channel which is already full is blocked until the value inside the channel
//is consumed by another goroutine.
package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	for i := 0; i < 5; i++ {
		ch <- "0example"
		fmt.Println("this is www.0example.com")
	}
}

//OUTPUT: An error called deadlock occurred as it is blocked after 3 values are queued to the channel.
//We can use select statement with default case to avoid this issue
