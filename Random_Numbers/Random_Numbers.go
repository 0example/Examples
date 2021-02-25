/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//In Go language, we can use package math/rand for generating pseudorandom numbers.
//It produces a deterministic sequence of values, each time program runs because it uses default shared Source. 
//If you need a different behaviour for each run, you have to use the seed function to initialize the default Source.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//Returns a integer as a random number between 0 to 10 ( 0 <= n < 10 )
	fmt.Println("Random Integer :",rand.Intn(10))

	// `rand.Float64` returns a `float64` `f`,
	// `0.0 <= f < 1.0`.
	//Returns a float as a random number in f[0.0,1.0)
	fmt.Println("Random float of range[0.0-1.0) :",rand.Float64())

	//We can return random floats in ranges
	//for example 10.00 <= f < 20.00
	fmt.Println("Random float of range[10.0-20.0) :",(rand.Float64()+1)*10)

	//To avoid getting same sequence of numbers in each run you have give a seed that changes every time.
	//You should use crypto/rand for random numbers when you
	// need to keep secret as this method is not safe.
	newSrc := rand.NewSource(time.Now().UnixNano())
	rand1 := rand.New(newSrc)
	fmt.Println("Random number with different seed each run :",rand1.Intn(10))

	//When you seed the source with a same number you will get same sequence of random numbers.
	src2 := rand.NewSource(42)
	r2 := rand.New(src2)
	fmt.Print("same seed : ")
	for i := 0; i <5; i++ {
		fmt.Print(r2.Intn(10),"\t")
	}
	fmt.Println()
	src3 := rand.NewSource(42)
	r3 := rand.New(src3)
	fmt.Print("same seed : ")
	for i := 0; i <5; i++ {
		fmt.Print(r3.Intn(10),"\t")
	}
}