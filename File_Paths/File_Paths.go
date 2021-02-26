/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//In Go language, package filepath implements utility routines for manipulating filename paths
//in a way compatible with the target operating system-defined file paths.
//Example : 
//		On Linux :	/home/John
//		On Windows :\Users\Rob

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	//joins the path elements into a hierarchical path, separating them with an OS specific Separator.
	path := filepath.Join("0example", "golang", "filepath.go")
	fmt.Println("path :", path)

	//You don't need to add OS specific separators manually.
	//If you add them join function will remove the superfluous separators.
	fmt.Println(filepath.Join("golang//package", "test.go"))
	fmt.Println(filepath.Join("golang/../package", "test.go"))

	//We can use `Dir` and `Base` to split a path to directory and file
	//Split function will split the path and return the directory and base at the same time
	fmt.Println("Dir of path(0example\\golang\\filepath.go):", filepath.Dir(path))
	fmt.Println("Base of path(0example\\golang\\filepath.go):", filepath.Base(path))
	d, b := filepath.Split(path)
	fmt.Println("Splitting of path(0example\\golang\\filepath.go):", d, b)

	//IsAbs function reports whether the path is absolute.
	fmt.Println(filepath.IsAbs("0example/golang.go"))
	fmt.Println(filepath.IsAbs("c://0example/golang.go"))

	filename := "golang.go"

	//Ext func returns the file name extension used by path.
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Get the file's name without the extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	//returns a relative path which is equivalent to target path
	//when joined to base path with an intervening separator.
	rel, err := filepath.Rel("ab/cd", "ab/cd/xy/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("ab/cd", "ab/ef/rs/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
