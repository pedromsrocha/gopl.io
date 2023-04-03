// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraw = make(chan struct {
	amount int
	ok     chan bool
})

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	ok := make(chan bool)
	withdraw <- struct {
		amount int
		ok     chan bool
	}{amount: amount,
		ok: ok}
	return <-ok
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case x := <-withdraw:
			if balance < x.amount {
				x.ok <- false
			}
			balance -= x.amount
			x.ok <- true
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
