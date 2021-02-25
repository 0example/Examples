/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//A waitgroup waits until the goroutines inside the waitgroup have successfully executed.
//A waitgroup contains three functions.
//	1: Add(int) -  Add adds integer value to the WaitGroup counter. 
//					If the counter becomes zero, all goroutines blocked on Wait are released. 
//					
//	2: Done()   - Done decreases the waitgroup counter by one.
//	3: Wait()   - Wait blocks until the WaitGroup counter is zero.


package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	links:=[]string{
		"http://0example.com",
		"https://golang.org",
		"https://play.golang.org/",
		"https://tour.golang.org/",
		"http://google.com",
	}
	for _,link:=range links{
		wg.Add(1)
		go checkLink(link,&wg)
	}
	wg.Wait()
}

func checkLink(link string, wg *sync.WaitGroup) {
	_, err :=http.Get(link)
	if err != nil {
		fmt.Println(link,"website is down..!")
		wg.Done()
		return
	}
	fmt.Println(link,"website is running...")
	wg.Done()
}

//First we create a waitgroup and add goroutines to it 
//When a goroutine is added it will increment the waitgroup counter. 
//We call function Add() on our WaitGroup to set the number of goroutines we want to wait for, 
//and then, we call function Done() at the end of function to indicate the end of goroutine's execution.         '
