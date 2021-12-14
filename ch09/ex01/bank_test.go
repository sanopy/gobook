package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	t.Run("Deposit", func(t *testing.T) {
		// Alice
		go func() {
			Deposit(200)
			fmt.Println("=", Balance())
			done <- struct{}{}
		}()

		// Bob
		go func() {
			Deposit(100)
			done <- struct{}{}
		}()

		// Wait for both transactions.
		<-done
		<-done

		if got, want := Balance(), 300; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
	})

	t.Run("Withdrawal", func(t *testing.T) {
		// Alice
		go func() {
			if ok := Withdraw(200); !ok {
				t.Errorf("Withdraw(200) failed")
			}
			fmt.Println("=", Balance())
			done <- struct{}{}
		}()

		// Bob
		go func() {
			if ok := Withdraw(100); !ok {
				t.Errorf("Withdraw(100) failed")
			}
			done <- struct{}{}
		}()

		// Carol
		go func() {
			if ok := Withdraw(500); ok {
				t.Errorf("Withdraw(500) succeeded over the amount of deposit")
			}
			done <- struct{}{}
		}()

		// Wait for all transactions.
		<-done
		<-done
		<-done

		if got, want := Balance(), 0; got != want {
			t.Errorf("Balance = %d, want %d", got, want)
		}
	})
}
