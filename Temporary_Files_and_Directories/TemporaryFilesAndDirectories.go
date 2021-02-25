/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Temporary files and directories are useful when we don''t need data after the program execution	

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// We can call ioutil.TempFile to create a temp file.
	// we can specify the directory we want to create these temp files
	//and define a pattern to name the temp files.
	file, err := ioutil.TempFile("img-dir", "golang-*.png")
	if err != nil {
		fmt.Println(err)
	}
	// Specify the file we want to be deleted on program close
	defer os.Remove(file.Name())

	fmt.Println(file.Name())

	//We can write some data into the file
	if _, err := file.Write([]byte("www.0example.com\n")); err != nil {
		fmt.Println(err)
	}

	// if the data cannot be read from the file
	// an error will be returned
	data, err := ioutil.ReadFile(file.Name())

	if err != nil {
		fmt.Println(err)
	}

	//printing the data as a string
	fmt.Print(string(data))

	//We can create a temporary directory
	dir, err := ioutil.TempDir("", "sampledir")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Temp dir name:", dir)

	defer os.RemoveAll(dir)

}