package main

import (
	accs "Alura_Banco/Accounts"
	"fmt"
	"reflect"
)

func main() {
	teste(1, 2, 3, 4, 5)
	account1 := accs.Account{"Gabriel Martins", 100, 10000, 500}
	account2 := accs.Account{"Eduardo", 200, 10000, 300}
	status, finalValue := account1.Transfer(-200, &account2)
	fmt.Println(status, finalValue)
	fmt.Println(account1)
	fmt.Println(account2)
}

func teste(a ...int) {
	fmt.Println(reflect.TypeOf(a))
}

func creatingSomeAccounts() {
	var account1 accs.Account = accs.Account{Name: "Gabriel Martins Teixeira Piagentini", BankNumber: 000, AccountNumber: 123456, Value: 00.000}
	var account2 accs.Account = accs.Account{"Gabriel Martins Teixeira Piagentini", 000, 123456, 00.000}
	var account3 accs.Account = accs.Account{Name: "Gabriel Martins Teixeira Piagentini", BankNumber: 123}
	var account4 accs.Account = *new(accs.Account)
	fmt.Println(account1)
	fmt.Println(account2)
	fmt.Println(account3)
	fmt.Println(account4)
}
