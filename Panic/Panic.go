/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Panic is a built in mechanism which handles error situations.
//When the panic function is called, it stops the ordinary flow of the program. 
//The panic function has a parameter: a message to show.
//You can use the panic function if an unexpected error arises in your go program
//that you don’t want or don’t know how to deal with.

package main

import "os"

func main() {
	_, err := os.Create("/root/0example")
	if err != nil {
		panic("Cannot create file")
	}
}

//File is needed to be created and if it cannot create the file 
//the panic will terminate the function a exit the program after giving a message.
