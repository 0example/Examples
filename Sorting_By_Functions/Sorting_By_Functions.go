/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//We can also implement our own sorting schema, by implementing the len,less,swap functions in the sort interface.
//For an example, if we wanted to sort some strings not in an alphabetical order instead considering their lengths 
//we can use the sort interface
//the length of the collection is returned by less function
//Less receives two integers which indicate indices from the collection
//After less method is called swap method makes the necessary changes.


package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByNameLen implements sort.Interface based on the Name field.
type ByNameLen []Person

func (b ByNameLen) Len() int {
	return len(b)
}
func (b ByNameLen) Less(i, j int) bool {
	return len(b[i].Name) < len(b[j].Name)
}
func (b ByNameLen) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	family := []Person{
		{"James", 15},
		{"Rob", 21},
		{"Dennis", 20},
	}

	sort.Sort(ByNameLen(family))
	fmt.Println(family)

	sort.Sort(ByAge(family))
	fmt.Println(family)
}
