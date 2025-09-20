package main

import (
	"github.com/jerberlin/examples-ch9bank1/bank"
)

func main() {
	balance := bank.Balance()
	println("Initial balance: ", balance)

	for i := 1; i <= 100; i++ {
		go func() {
			// add some fixed amount
			amount := 1
			bank.Deposit(amount)
		}()
	}

	// check final balance: should be 100
	balance = bank.Balance()
	println("Final balance: ", balance)
}
