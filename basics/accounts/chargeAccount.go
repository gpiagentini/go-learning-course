package accs

import clients "Alura_Banco/Clients"

type Account struct {
	Client        clients.Client
	BankNumber    int
	AccountNumber int
	value         float64
}

func (a *Account) Withdraw(value float64) (string, float64) {
	canWithdraw := value > 0 && value <= a.value
	if canWithdraw {
		a.value -= value
		return "Withdraw successfuly executed!", a.value
	} else {
		return "Cannot withdraw!", a.value
	}
}

func (a *Account) Deposit(value float64) (string, float64) {
	canDeposit := value > 0
	if canDeposit {
		a.value += value
		return "Deposit successfuly executed!", a.value
	} else {
		return "Cannot withdraw!", a.value
	}
}

func (a *Account) Transfer(value float64, finalAccount *Account) (string, float64) {
	canTransfer := value > 0 && value <= a.value
	if canTransfer {
		a.value -= value
		finalAccount.value += value
		return "Transfer successfult executed!", a.value
	} else {
		return "Cannot transfer!", a.value
	}
}

func (a *Account) GetValue() float64 {
	return a.value
}
