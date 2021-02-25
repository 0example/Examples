/*	Copyright 2021 www.0example.com, powered by Wixis360  */


//We can implement the predeclared error interface to define a customer error type.

package main

import (
	"fmt"
	"log"
)

//The bank type is defined.
type bank struct {
	balance float64
	amount  float64
	msg     string
}


//The error interface is implemented.
func (e *bank) Error() string {
	return fmt.Sprintf("amount :%0.2f -- balance :%0.2f -- msg: %s", e.balance, e.amount, e.msg)
}

//the enterAmount function checks whether the balance and the amount satisfy their requirements
//and if the amount can be withdrawn it will return the remaining balance.
//If they cannot satisfy the requirements an error will be generated.
func enterAmount(bal float64, amnt float64) (string, error) {
	if bal <= 0 || amnt > bal || amnt <= 0 {
		return "", &bank{bal, amnt, "Invalid amount or balance"}
	}
	return fmt.Sprintf("remaining balance : %0.2f ", bal-amnt), nil
}

func main() {

	var balance, amount = 1000.00, 200.00

	msg, err := enterAmount(balance, amount)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)

	balance, amount = 150.00, 2000.00
	msg, err = enterAmount(balance, amount)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
