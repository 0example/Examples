/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//Even though Go language doesn't have classses, we can define methods on types.
//A method is a function which contains a receiver argument.

package main

import "fmt"

// Student struct
type student struct {
	id string
	name string
	email string
	contactNumber int64
}

//Receivers in methods can be pointers or value receiver types.
func (s student) desc() {
	fmt.Println("Student ID : ", s.id)
	fmt.Println("Student Name : ", s.name)
	fmt.Println("Email : ", s.email)
	fmt.Println("Contact Number: ", s.contactNumber)
}

func (s *student) changeEmail(em string){
	s.email=em
}

func main() {

	// Initializing student struct
	res := student{
		id: "st001",
		name:	 "James",
		email: "james@gmail.com",
		contactNumber: 0236424203,
	}

	// Calling methods
	res.desc()
	res.changeEmail("james.arthur@gmail.com")
	fmt.Print(res.email)

}
