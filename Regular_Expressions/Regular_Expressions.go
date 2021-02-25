/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//A regular expression is a sequence of characters use  for searching  and replacing text 
//and more advanced text manipulation.
//In go language, package regexp implements regular expression search.


package main

import (
	"fmt"
	"regexp"
)

func main() {
	//find string Method
	//Returns a string after matching the given pattern.
	//If there is no match a null value is returned.
	regex := regexp.MustCompile("le$")
	fmt.Println(regex.FindString("0example"))
	fmt.Println(regex.FindString("golang"))

	//Find String Indexes method
	//Returns starting and ending indexes of the matching string
	//If there is no match a null value is returned.
	regex2 := regexp.MustCompile("ex")
	fmt.Println(regex2.FindStringIndex("0example"))

	//find String Submatch method
	//Returns a string that matches the regex pattern
	//If there is no match a null value is returned
	regex3 := regexp.MustCompile("0([a-z]+)e")
	fmt.Println(regex3.FindStringSubmatch("0example"))

	//find String Submatch Index method
	//Returns starting and ending indexes matches and submatches
	//If there is no match a null value is returned
	regex4 := regexp.MustCompile("0([a-z]+)e")
	fmt.Println(regex4.FindStringSubmatchIndex("0example"))

	//findAllstring method
	//returns a multiple matches of a pattern within a string.
	//It returns values as a slice
	regex5 := regexp.MustCompile(`e.`)
	fmt.Println(regex5.FindAllString("elephent",-1))

}