/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Golang select statement is similar to the switch statement, 
//which lets a goroutine wait on multiple channel operations.

//The select statement blocks as a whole unit, until one of the operations finished 

package main

import (
	"fmt"
	"time"
)

func calling01(c1 chan string) {
	time.Sleep(10 * time.Second) //sleeps for 10 seconds
	c1 <- "Calling method 1"
}

func calling02(c2 chan string) {
	time.Sleep(5 * time.Second) //sleeps for 5 seconds
	c2 <- "Calling method 2"
}

func main() {

	chn1 := make(chan string, 1)
	chn2 := make(chan string, 1)

	go calling01(chn1)
	go calling02(chn2)
	
	//blocks until there's data available on chn1 or chn2
	select {
	case msg1 := <-chn1:
		fmt.Println(msg1)  //
	case msg2 := <-chn2:
		fmt.Println(msg2)  // receives value  after 5 seconds.
	}

}

//Values will return to the two channels after passing the relevent time amount spent in the relevent method. 
//The output depends on the time taken to execute the relevent methods.
//Therefore, after 5 seconds