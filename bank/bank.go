package bank

var (
	deposit    func(int)
	getBalance func() int
)

// Deposit sends an amount to be added to the balance
func Deposit(amount int) { deposit(amount) } // call the closure buit at init

// Balance returns the current balance
func Balance() int { return getBalance() } // call the closure built at init

// teller is the monitor goroutine, confining var balance.
func teller() (func(int), func() int) {
	var balance int // balance will be captured and hence be persisted

	write := func(amount int) { balance += amount }
	read := func() int { return balance }

	return write, read
}

func init() {
	deposit, getBalance = teller() // build the closures through teller()
}
