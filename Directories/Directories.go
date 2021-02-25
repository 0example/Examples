/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//A directory is a unit in a computer''s file system for storing and locating files.


package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func checkForErr(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	// function Mkdir creates a new directory with the specified name
	//in the current working directory.
	err := os.Mkdir("temp-dir", 0755)
	checkForErr(err)

	//function os.RemoveAll removes path and any children it contains
	defer os.RemoveAll("temp-dir")

	//This function is used to create a new Empty File with a specified name and permission bits.
	createEmptyFile := func(filename string) {
		d := []byte("")
		checkForErr(ioutil.WriteFile(filename, d, 0644))
	}

	createEmptyFile("temp-dir/myfile.txt")

	//function `MkdirAll` creates a directory, alone with any necessary parents
	//If the directory already exists, MkdirAll does nothing and returns nil.
	err = os.MkdirAll("temp-dir/parent-dir/child-dir", 0755)
	checkForErr(err)

	createEmptyFile("temp-dir/parent-dir/myfile2.txt")
	createEmptyFile("temp-dir/parent-dir/myfile3.txt")
	createEmptyFile("temp-dir/parent-dir/child-dir/myfile4.txt")

	// `ReadDir` lists directory contents, returning a
	// slice of `os.FileInfo` objects.

	//Returns a list of directory entries sorted by filename after reading the directory.
	dirList, err := ioutil.ReadDir("temp-dir/parent-dir")
	checkForErr(err)

	fmt.Println("Listing temp-dir/parent-dir")
	for _, entry := range dirList {
		fmt.Println("  ", entry.Name(), entry.IsDir())
	}

	//function `Chdir` changes the current working directory to the named directory.
	//It is similar to `cd`
	err = os.Chdir("temp-dir/parent-dir/child-dir")
	checkForErr(err)

	//changes the current working directory to where we started
	err = os.Chdir("../../..")
	checkForErr(err)

	//Function `Walk` walks a file tree calling a function of type filepath.WalkFunc
	//for each file or directory in the tree, including the root.
	fmt.Println("Visiting temp-dir")
	err = filepath.Walk("temp-dir", visit)
}

//`visit` is called for every file or directory found recursively by `filepath.Walk`.
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}