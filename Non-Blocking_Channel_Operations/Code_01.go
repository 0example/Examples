/*	Copyright 2021 www.0example.com, powered by Wixis360  */

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