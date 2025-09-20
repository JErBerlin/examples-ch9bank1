package main

import (
	"github.com/jerberlin/examples-ch9bank1/bank"
)

func main() {
	balance := bank.Balance()
	println("Initial balance: ", balance)

	// add some fixed amount
	amount := 100
	bank.Deposit(amount)
	println("Add the amount: ", amount)

	// check final balance
	balance = bank.Balance()
	println("Final balance: ", balance)
}
