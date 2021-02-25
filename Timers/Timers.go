/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Timer is a built in feature in go that can be used to fire some task after a period of time in the future.
//NewTimer function creates a new Timer that will send the current time on its channel
// after duration which was passed as an arguement.
//NewTimer(d).C is equivalent to the time.After(d) function.

//If the timer is no longer needed we can use NewTimer instead and call Timer.Stop()
//It will prevent the timer from firing and return true if the call stop the timer.

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		t1 := time.After(time.Second)
		<-t1
		fmt.Println("timer ONE executed.. :)")
	}()

	go timer2()

	t3 := time.NewTimer(3*time.Second)
	go func() {

		<-t3.C
		fmt.Println("time THREE executed.. :)")
	}()


	if t3.Stop(){
		fmt.Println("timer THREE terminated..!")
	}


	time.Sleep(4 * time.Second)
}
func timer2() {
	t2 := time.NewTimer(2*time.Second)
	<-t2.C
	fmt.Println("timer TWO executed.. :)")
}