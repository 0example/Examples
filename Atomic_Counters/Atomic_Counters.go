/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//We communicate with channels to manage state in go.
//Atomic counters also can be utilized to manage state and avoid race condtinos in go.
//Atomic counters can be accessed by multiple goroutines.
//They are more primitive than other synchronization techniques and implemented directly at hardware level. 
//Package `atomic` provides low-level atomic memory primitives useful for implementing synchronization algorithms.


package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
func main() {

	// A WaitGroup will help us wait for all goroutines
	// to finish their work.

	//Using sync.WaitGroup to wait for all the goroutines to finish
	var wg sync.WaitGroup

	// Declaring atomic variable
	var counter int64

	//will start 7 goroutines
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter,1)  //counter = counter + 1
			fmt.Println(atomic.LoadInt64(&counter)) // read the counter value
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("final value of the counter :",counter)
}