/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Go language supports for opening and writing files on the hard drive.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//This will create a file and overwrite the file if it already exists
	file, err := os.Create("myfile.txt")

	if err != nil {
		log.Fatal(err)
	}

	//Defer a close after opening a file.
	defer file.Close()

	//The File.WriteString function writes a string to a file
	//and returns the number of bytes written and an error if an error occurred.
	num1, err2 := file.WriteString("alpha\n")
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of bytes written: %d\n",num1)


	bs1 := []byte("beta\n")

	//The File.Write function writes byte slices to a file and
	//and returns the number of bytes written and an error if an error occurred.
	num2, err3 := file.Write(bs1)

	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Printf("Number of bytes written: %d\n",num2)

	bs2 := []byte("omega\n")

	//The File.WriteAt function writes byte slices to a file
	//starting at the specified byte offset.
	i := int64(len(bs1)+len(bs2))

	_, err4 := file.WriteAt(bs2, i)

	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Println("Completed.")
}
