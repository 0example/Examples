/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Tickers can execute a certain peace of code repeatedly on a regular interval of time. 

//	Timers - use for one time task firing in a future time.
//	Tickers -use for repeated tasks at regular time intervals.




package main

import "fmt"
import "time"

func main() {
	go ticking()
	time.Sleep(2*time.Second)
}

func ticking() {
	c := time.Tick(550*time.Millisecond)
	for t:= range c{
		fmt.Println("Tick at...",t)
	}
}

//As shown in the example you can use tick function combined with range to repeat some task.
//Tick is a wrapper for NewTicker function which provides access to the ticking channel. 
//Tick is useful for clients that have no need to shut down the Ticker. 
//But ticker cannot be recovered by garbage collector as there is no way to shut down the ticker.

