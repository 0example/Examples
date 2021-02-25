/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//In Go Language ,we can use the ReadFile() function of the ioutil library to read data from a file.
//We can read text and binary files using this function.
//When there is large files to read we can convert them into byte slices and easily read.


package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Use to help checking for errors when reading files.
func checkForErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Reading file into memory
	//Here, whole file is read into a string.
	//Even though this method is pretty straightforward, it should not be used for larger files
	st, err := ioutil.ReadFile("path/student.txt")
	checkForErr(err)
	fmt.Print(string(st))


	//Open function opens the named file for reading.
	//methods on the returned os.file can be used for reading
	file, err := os.Open("path/student.txt")
	checkForErr(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.

	//reads bytes from the File up to length of the byte slice passed as an argument.
	//It returns the number of bytes read and any error encountered
	bs1 := make([]byte, 14)
	count, err := file.Read(bs1)
	checkForErr(err)
	fmt.Println("No of bytes read :",count)
	fmt.Println("Data read from the file :",string(bs1[:count]))


	// Using seek function we can point to a specified location in the file
	// and read from there
	//It returns the new offset and an any error encountered.
	ofs2, err := file.Seek(6, 0)
	checkForErr(err)
	bs2 := make([]byte, 6)
	count2, err := file.Read(bs2)
	checkForErr(err)
	fmt.Printf("%d bytes from position %d : %v\n ", count2, ofs2, string(bs2[:count2]))

	//ReadAtLeast reads from the file until it has read at least, specified amount of bytes.
	//It returns the number of bytes copied and an error if fewer bytes were read.
	ofs3, err := file.Seek(6, 0)
	checkForErr(err)
	bs3 := make([]byte, 3)
	count3, err := io.ReadAtLeast(file, bs3, 3)
	checkForErr(err)
	fmt.Printf("%d bytes from %d to : %s\n", count3, ofs3, string(bs3))

	//When you don't need to read the entire file but only a specific part
	//we can use package bufio to read files in chunks of data.
	readr4 := bufio.NewReader(file)
	bs4, err := readr4.Peek(3)
	checkForErr(err)
	fmt.Printf("3 bytes: %s\n", string(bs4))

	//You should close the file when you complete reading
	file.Close()
}