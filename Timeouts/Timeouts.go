/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//We always try to make Real life applications time sensitive. 
//Therefore, Timeouts are necessary to ensure that tasks that run overdue 
//do not consume the resources that might be needed by other components of our application. 
//We can easily implement timeouts by using channels and select.


package main

import (
	"fmt"
	"time"
)
func main() {
	c1 := make(chan string, 1)
	go funcABC(c1)

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("ABC timeout.! :(")
	}


	c2 := make(chan string, 1)
	go funcXYZ(c2)

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("XYZ timeout.! :(")
	}
}

func funcABC(c1 chan string){
	time.Sleep(2* time.Second)
	c1 <- "ABC running.. :)"
}
func funcXYZ(c2 chan string) {
	time.Sleep(6* time.Second)
	c2 <- "XYZ running.. :)"
}

//In the example above,Two functions funcABC and funcXYZ are executed in goroutines 
//which take 2s and 6s to complete respectively.
//Select statements are implementing timeouts res := <-c1 and res := <-c2 await for results to return.
//and <-time.After(3 * time.Second) awaits to send the timeout message 6 seconds.
//If the channel recieves a value before the timeout time then the result is printed( ex: funcABC)
//If it is timeout before receiving a value then the timeout value is printed (ex: funcXYZ)
