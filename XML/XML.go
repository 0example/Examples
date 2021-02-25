/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Go has built-in support for XML by implementing the encoding.xml package.

package main

import (
	"encoding/xml"
	"fmt"
)

// Here, Student maps to XML and field tags provide element paths for the encoder and decoder.
// `XMLName` field name set the name of the XML element which represents the struct;
// A field with tag "name,attr" becomes an attribute with the given name in the XML element.
type Student struct {
	XMLName xml.Name `xml:"student"`
	Id      string      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Email    string   `xml:"email"`
	Contact    int   `xml:"contact"`
}

func (s Student) String() string {
	return fmt.Sprintf("Student id=%v, name=%v, email=%v, contact:=%v",
		s.Id, s.Name, s.Email, s.Contact)
}

func main() {
	james := &Student{Id: "S001", Name: "James",Email: "james@gmail.com",Contact: 12234344}

	// Receiving an XML representing student
	xmlb, _ := xml.MarshalIndent(james, " ", "  ")
	fmt.Println(string(xmlb))

	// Adding a generic XML header to the output
	fmt.Println(xml.Header + string(xmlb))

	//Useing Unmarshal function to parse byte array with XML into a data structure.
	//An error will be returned if the xml can not be formed properly or xml cannot be mapped to data structure .
	var s2 Student
	if err := xml.Unmarshal(xmlb, &s2); err != nil {
		panic(err)
	}
	fmt.Println(s2)

	arthur := &Student{Id: "S002", Name: "Arthur",Email: "arthur@gmail.com",Contact: 22234344}

	//Encoder identify to nest students under <parent><child> by `parent>child>student` field tag
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Students  []*Student `xml:"parent>child>student"`
	}

	nesting := &Nesting{}
	nesting.Students = []*Student{james, arthur}

	xmlb, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(xmlb))
}