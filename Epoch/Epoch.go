/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//In Go language, we can use time.Now() function to get current time and  
//Unix() function to convert it into an epoch timestamp

package main

import (
	"fmt"
	"time"
)

func main() {

	//In Go, we can use the time.Now() function to get current time
	currentTime := time.Now()

	//Using, Unix() function to convert the time into an epoch timestamp(In seconds):
	epochSec := currentTime.Unix()

	fmt.Println("current time :",currentTime)

	fmt.Println("epoch time stamp in seconds :",epochSec)

	//Converting unix seconds into time
	timeSec := time.Unix(epochSec, 0)
	fmt.Println("Convert epoch to relevent time :",timeSec)

	//Converting to unix epoch in nanoseconds
	epochNano := currentTime.UnixNano()
	fmt.Println("epoch timestamp in nanoseconds :",epochNano)

	//Converting unix nonoseconds into time
	timeNano := time.Unix(0, epochNano)
	fmt.Println("convert epoch to relevent time :",timeNano)

}