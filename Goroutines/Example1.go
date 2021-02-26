/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Concurrency
//	-Concurrency is the ability of independently executing multiple tasks at the same time.
//	-But it is not parallelism.
//	-Golang has 2 built-in concurrency constructs: goroutines & channels

//Goroutines
//	-Goroutine is an function that executes concurrently, managed by go runtime.
//	-You can see it as a lightweight thread.
//	-To invoke goroutine you should add `go` keyword before calling the function.

package main

import (
	"fmt"
	"time"
)

func main() {
	go counting()   // Using "go" keyword to invoke the new goroutine
	time.Sleep(5*time.Second)    //sleep main goroutine for 5 seconds
}

func counting() {
	for i := 0; ; i++ {
		fmt.Println("counting :", i)
		time.Sleep(500* time.Millisecond) //slowing down the loop
	}
}

//Here, after running the above code when the main function executes the main goroutine starts running
//then due to Sleep method the main Goroutine sleeps for 3 seconds in between 3 seconds,
//the new Goroutine executes displaying the counting numbers on the screen.
//When main method returns after the duration given in the sleep method,
//the program exits and takes the counting function down with it.