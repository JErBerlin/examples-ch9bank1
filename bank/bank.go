package bank

var deposits = make(chan int) // to send money to the deposit
var balances = make(chan int) // to receive the balance of the deposit

// Deposit sends an amount to be added to the balance
func Deposit(amount int) { deposits <- amount }

// Balance returns the current balance
func Balance() int { return <-balances }

// teller is the monitor goroutine, confining var balance.
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start monitor goroutine

}
