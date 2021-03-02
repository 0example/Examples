/*	Copyright 2021 www.0example.com, powered by Wixis360  */

//Environment variable is a dynamic key-value pair on the operating system
//and it is used to configure running processes' behaviours on the OS.
//The `go` command uses the default setting, if an environment variable is unset.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//os.Setenv() - Sets the value of an environment variable named by the key.
	os.Setenv("WIXIS_USER", "gopher")

	//os.Getenv() - Gets the value environment variable named by the key.
	fmt.Println("WIXIS_USER:", os.Getenv("WIXIS_USER"))

	//An empty string will be returned, if the variable named by the key
	//is not present in the environment.
	fmt.Println("WIXIS_HOST:", os.Getenv("WIXIS_HOST"))

	//os.Environ() - Returns a slice of the string of all environment variables in the form of key=value.
	//We can separate them into key and value pairs with strings.SplitN.
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		fmt.Println(pair[0])
	}
}