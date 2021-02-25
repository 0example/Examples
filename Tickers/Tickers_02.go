/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Tickers can be stopped like timers using stop() function when you are using newTicker function.	


package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500* time.Millisecond) // set to send values after every 500ms to the channel
	done := make(chan bool)

	go func() {
		time.Sleep(2400 * time.Millisecond)
		ticker.Stop()  // ticker is stopped after 2400ms. channel won't receive any values after this.
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("ticker stopped..!")
			return
		case t := <-ticker.C: // recieves values after every 500ms.
			fmt.Println("time : ", t)
		}
	}

}

//The NewTicker() function returns a new Ticker containing a channel 
//which send time according to the duration argument. 
//The duration must be greater than zero, if not, the ticker will panic.
//Unlike NewTicker, Tick will return nil if duration is <= 0.
