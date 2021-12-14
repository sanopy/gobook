// Package bank provides a concurrency-safe bank with one account.
package bank

var (
	deposits    = make(chan int)        // send amount to deposit
	balances    = make(chan int)        // receive balance
	withdrawals = make(chan withdrawal) // send amount of withdraw
)

type withdrawal struct {
	amount int
	result chan bool
}

func NewWithdrawal(amount int) withdrawal {
	return withdrawal{amount, make(chan bool)}
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	w := NewWithdrawal(amount)
	withdrawals <- w
	return <-w.result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdrawals:
			if w.amount <= balance {
				balance -= w.amount
				w.result <- true
			} else {
				w.result <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
