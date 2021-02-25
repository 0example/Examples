/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//by using the default case we can avoid the deadlock error.
package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	for i := 0; i < 5; i++ {
		select {
		case ch <- "0xeample":
			fmt.Println("message sent to channel successfully :)")
		default:
			fmt.Println("Channel is full. :(")
		}
	}
}
//OUTPUT:
//message sent to channel successfully :)
//message sent to channel successfully :)
//message sent to channel successfully :)
//Channel is full. :(
//Channel is full. :(