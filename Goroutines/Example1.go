/*	Copyright 2021 www.0example.com, powered by Wixis360  */

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