package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t,err)
		assertBalance(t, wallet, Bitcoin(5))
	})
	
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		initialBalance := Bitcoin(20)
		wallet.Deposit(initialBalance)
		err := wallet.Withdraw(Bitcoin(25))
		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientFund)
	})
}

func assertBalance (t testing.TB, wallet Wallet, want Bitcoin) {
	balance := wallet.Balance()
	if balance != want {
		t.Errorf("Got %s want %s", balance, want)
	}
}

func assertError (t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected an error but did not recieve any")
	}
	if err.Error() != want {
		t.Errorf("Wanted error '%s' but got '%s'", want, err.Error())
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Expected no error but received error ", err.Error())
	}
}