/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//In Go Language, Range can be used to iterate over strings, arrays, slices, maps and channels.
//For each iteration, two values are returned. The first is the index, 
//and the second is a copy of the element at that index.

package main

import "fmt"

func main() {
	//Using range to iterate over an Array
	arr := [3]string{"golang", "0example", "wixis360"}
	for i, str := range arr {
		fmt.Printf("Index %v :\tvalue :%v\n", i, str)

	}

	//Using range to iterate over a slice
	sl := []int{245, 554, 632, 574, 389}
	sum := 0
	for _, num := range sl {
		sum += num
	}
	fmt.Print("sum :", sum)

	//Using range to iterate over a string.
	str := "0example"
	for i, e := range str {
		fmt.Printf("%d %c\n", i, e)
	}

	//Using range to iterate over a map.
	m := map[string]string{
		"Sri Lanka":      "LKR",
		"United States":  "USD",
		"United Kingdom": "Sterling Pound",
	}
	for key, val := range m {
		fmt.Println(key, " - ", val)
	}

	for key := range m {
		fmt.Println(key, " - ", m[key])
	}

	//Using range to iterate over a channel
	chn := make(chan string, 2)

	chn <- "on"
	chn <- "off"
	close(chn)

	for n := range chn {
		fmt.Println(n)
	}
}
