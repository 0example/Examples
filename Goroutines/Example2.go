/*	Copyright 2021 www.0example.com, powered by Wixis360  */

package main

import (
	"fmt"
	"time"
)

func main() {

	go print("GoRoutine")
	print("Direct")

	go func(str string) {
		fmt.Println(str)
	}("Annonymous Func")

	fmt.Println("Back to Main GoRoutine")

	time.Sleep(3*time.Second) //pause the main goroutine for 3 seconds
}

func print(str string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(str,i)
	}
}
//Invoking print function in a goroutine by using go keyword. 
//secondly, calling print function the usual way running it synchronously.
//We can use goroutine for an annonymous function as well.
//After executing the code we can see that the functions are running asynchronously in seperate goroutines.