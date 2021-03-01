/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//A regular expression is a sequence of characters use  for searching  and replacing text
//and more advanced text manipulation.
//In go language, package regexp implements regular expression search.

package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	//Tests whether the string contains any match of the given pattern and returns a boolean value.
	//For more complicated queries, we should compile a regular expression to create a Regexp object
	//and there are two methods to compile.
	isAMatch, _ := regexp.MatchString("go", "gopher")
	fmt.Println(isAMatch)

	//Tests whether the byte slice contains any match of the regular expression pattern.
	isMatch2, err := regexp.Match(`go.*`, []byte(`gopher`))
	fmt.Println(isMatch2, err)

	//Compile function parses the regular expression and if it was successful,
	//It returns a Regexp object that can be used to match against text.
	//if there was an error while parsing it returns the error.
	r1, err := regexp.Compile(".....ple?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//FindString Method
	//It returns a substring holding the text of the leftmost match by matching the given pattern.
	//If there is no match an empty string is returned as the return value.
	fmt.Println(r1.FindString("0example pineapple"))
	fmt.Println(r1.FindString("pineapple 0example"))
	fmt.Println(r1.FindString("www.golang.org"))

	//MustCompile function is similar to Compile function but it panics
	//instead of returning an error, if the expression cannot be parsed.
	//We can use mustcompile function when creating global variables with regular expressions because,
	//it simplifies safe initialization of global variables holding compiled regular expressions.
	r2 := regexp.MustCompile("pl")

	//FindStringIndex Method
	//It returns a two-element slice of integers defining the location of the leftmost match
	//with starting and ending indexes of the regular expression
	//If there is no match nil value is returned.
	fmt.Println(r2.FindStringIndex("0example pineapple"))

	//FindStringSubmatch Method
	//It returns a slice of strings containing the text of the leftmost match of the regex pattern.
	//If there is no match nil value is returned
	r3 := regexp.MustCompile("0([a-z]+)e")
	fmt.Println(r3.FindStringSubmatch("0example pineapple"))

	//FindStringSubmatchIndex Method
	//It returns a slice containing starting and ending index pairs of matches and submatches
	//If there is no match a null value is returned
	r4 := regexp.MustCompile("0([a-z]+)e")
	fmt.Println(r4.FindStringSubmatchIndex("0example pineapple"))

	//FindAllString Method
	//It is the 'All' version of FindString method.
	//It returns slice of multiple matches of the expression.
	r5 := regexp.MustCompile(`e.`)
	fmt.Println(r5.FindAllString("0example pineapple", -1))

	r6 := regexp.MustCompile(`a(p*)l`)
	//Returns a copy of string after replacing string with the specified regular expression.
	fmt.Println(r6.ReplaceAllString("apple", "B"))

}
