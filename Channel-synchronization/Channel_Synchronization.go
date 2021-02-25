/*	Copyright 2021 www.0example.com, powered by Wixis360  */

package main

import (
	"fmt"
	"time"
)

func task1(ok chan bool) {
	fmt.Print("Doing task 1...")
	time.Sleep(2*time.Second)		//Sleeps for two seconds
	fmt.Println("finished")

	ok <- true		//Sends ok when the task1 is finished
}
func task2(ok chan bool) {
	fmt.Print("Doing task 2...")
	time.Sleep(time.Second) //sleeps for a second.
	fmt.Println("finished")

	ok <- true 		//Sends true when the task2 is finished
}

func main() {

	ok := make(chan bool)

	go task1(ok)
	<-ok   //block the goroutine until receive a value
	go task2(ok)
	<-ok //block the goroutine until receive a value
}
