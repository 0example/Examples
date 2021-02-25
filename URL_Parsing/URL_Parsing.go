/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Go language supports for url parsing. 
//URL has a scheme, authentication info, host, port, path, query params, query patterns and query fragment. 
//Package url parses URLs and implements query escaping.


package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	rawURL :="http://user:pass@0example.com:8000/docs/books/tutorial/index?key2=networking#DOWNLOADING"
	p := fmt.Println

	//Parse the URL and if there is an error returns an error message.
	URL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	//Returns Scheme of the URL
	p("Scheme :",URL.Scheme)

	//Returns username and password
	p("User and Password :",URL.User)

	//Returns Username
	p("Username : ",URL.User.Username())

	//Returns user password
	pass, _ := URL.User.Password()
	p("Password :",pass)

	//Returns 	host and port
	p("Host and Port :",URL.Host)

	//Splits host name and port
	host, port, _ := net.SplitHostPort(URL.Host)
	p("Host :",host)
	p("Port :",port)

	//Returns path
	p("Path :",URL.Path)

	//Prints fragment path value
	p("Fragment :",URL.Fragment)

	//Returns Query param name and value
	p("Query param name and value :",URL.RawQuery)

	//Parse query param into map
	m, _ := url.ParseQuery(URL.RawQuery)
	p("Param Map :",m)

	//Prints the relevent key value
	p("Key Value :",m["key2"][0])
}