/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Mutex or Mutual Exclusion lock is a is a synchronization technique which can be used to 
//synchronize access to state and safely access data across multiple goroutines. 
//We can define a block of code to be executed in a mutex by surrounding it with a call to Lock and Unlock 
//The zero value for a Mutex is an unlocked mutex.
//A Mutex must not be copied after first use.


package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex  // synchronize access to the state

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Counter : ",counter)
}

//Here, codes present between to Lock and Unlock will be executed by only one Goroutine at a time.
func incrementor(str string){
	for i := 0; i < 10; i++ {
		time.Sleep(200*time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(str,i,"counter:",counter)
		mutex.Unlock()
	}
	wg.Done()
}