/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Package strings provides a number of built-in string functions to manipulate UTF-8 encoded strings.
//Here are some predifined functions in springs package.


package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println


func main() {

	//Contains reports whether substr is within s.
	p("Contains:  ", s.Contains("0example", "0"))

	//counts the number of non-overlapping instances of substr in s.
	p("Count:     ", s.Count("0example", "p"))

	//HasPrefix tests whether the string s begins with prefix.
	p("HasPrefix: ", s.HasPrefix("0example", "0ex"))

	//tests whether the string s ends with suffix.
	p("HasSuffix: ", s.HasSuffix("0example", "le"))

	//returns the index of the first instance of substr in s, or -1 if substr is not present in s.
	p("Index:     ", s.Index("0example", "e"))

	//concatenates the elements of its first argument to create a single string.
	//The separator string sep is placed between elements in the resulting string.
	p("Join:      ", s.Join([]string{"0example", "com"}, "."))

	//returns a new string consisting of count copies of the string s.
	p("Repeat:    ", s.Repeat("0", 5))

	//returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	p("Replace:   ", s.Replace("0example", "e", "E", -1))

	//returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	p("Replace:   ", s.Replace("0example", "e", "E", 1))

	//Split slices s into all substrings separated by sep and
	//returns a slice of the substrings between those separators.
	p("Split:     ", s.Split("www.0example.com", "."))

	//returns s with all Unicode letters mapped to their lower case.
	p("ToLower:   ", s.ToLower("0EXAMPLE"))

	//returns s with all Unicode letters mapped to their upper case.
	p("ToUpper:   ", s.ToUpper("0example"))
}