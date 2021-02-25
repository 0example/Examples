/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Go language supports for opening and writing files on the hard drive.


package main

import (
	"log"
	"os"
)

func main() {
	// This will overwrite the file if it already exists
	file, err := os.Create("myfile.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	//The File.WriteString function writes a string to a file.
	_, err2 := file.WriteString("alpha\n")
	if err2 != nil {
		log.Fatal(err)
	}

	//The File.Write function can write byte slices to a file.
	bs1 := []byte("beta\n")

	_, err3 := file.Write(bs1)

	if err3 != nil {
		log.Fatal(err3)
	}

	//The File.WriteAt function writes byte slices to a file
	//starting at the specified byte offset.
	bs2 := []byte("omega\n")

	i := int64(len(bs1)+len(bs2))

	_, err4 := file.WriteAt(bs2, i)

	if err4 != nil {
		log.Fatal(err4)
	}
	//
}