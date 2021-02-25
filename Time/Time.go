/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//In Go language, package time provides functionality for measuring and displaying time.
//Go supports for various types of time manipulations.
//we can built a time object by using Date function provided in the time package.
//There are functions like year(), month(), day() in the time package.


package main

import (
	"fmt"
	"time"
)
func main() {

	currentTime := time.Now()

	//Here are various time formats

	fmt.Println("Current Time : ", currentTime.String())

	fmt.Println("MM-DD-YYYY\t:", currentTime.Format("01-02-2006"))

	fmt.Println("YYYY.MM.DD\t:", currentTime.Format("2006.01.02 15:04:05"))

	fmt.Println("YYYY:MM:DD {Character} :", currentTime.Format("2006:01:02"))

	fmt.Println("YYYY-MM-DD hh:mm:ss:", currentTime.Format("2006-01-02 15:04:05"))

	fmt.Println("Time with MicroSeconds:", currentTime.Format("2006-01-02 15:04:05.000000"))

	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))

	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))

	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))

	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))

	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))

	//Can initialize a time struct by providing year,month,day, hour, min, sec, nsec, and location in UTC.
	pastDateTime := time.Date(
		2012, 12, 12, 12, 12, 12, 1212121212, time.UTC)
	fmt.Println("Past Date Time :",pastDateTime)

	//Can extract details from the time struct built.
	fmt.Println("Year of pastDateTime :",pastDateTime.Year())
	fmt.Println("Month of pastDateTime :",pastDateTime.Month())
	fmt.Println("Date of pastDateTime :",pastDateTime.Day())
	fmt.Println("Hour of pastDateTime :",pastDateTime.Hour())

	//We can compare two times and check whether which happened before or after or at the same time
	fmt.Println("Whether pastDateTime before current time :",pastDateTime.Before(currentTime))
	fmt.Println("Whether pastDateTime after current time :",pastDateTime.After(currentTime))
	fmt.Println("Whether pastDateTime equals to current time :",pastDateTime.Equal(currentTime))

	//We can get the duration between  a time interval
	dur := currentTime.Sub(pastDateTime)
	fmt.Println(dur)

	//We can calculate the duration in different time units.
	fmt.Println("Duration in hours : ",dur.Hours())
	fmt.Println("Duration in Minutes : ",dur.Minutes())
	fmt.Println("Duration in Seconds : ",dur.Seconds())


}