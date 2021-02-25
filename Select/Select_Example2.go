/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//The default case will be executed if all other cases are blocked.


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

	select {
	case msg1 := <-chn1:
		fmt.Println(msg1)
	case msg2 := <-chn2:
		fmt.Println(msg2)
	default:
		fmt.Println("www.0example.com")
	}
}

//Default case is used to protect the select statement from blocking.
//If the Select has no default case, the select statement waits until at least one 
//case can be executed.
//When there is a default case, it is executed if none of other cases are ready to proceed.