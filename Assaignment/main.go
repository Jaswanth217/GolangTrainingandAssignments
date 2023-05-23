package main

import (
	"GolangTrainingandAssignments/Assaignment/App"
	"fmt"
)

func main() {
	//Add Accounts

	myAcc := App.Account{"jaswanth", "Pelaprolu", 1, 10000}
	lstAccount := App.AddAccount(myAcc)
	fmt.Println(lstAccount)

	myAcc1 := App.Account{"Sri", "Sainath", 2, 20000}
	lstAccount = App.AddAccount(myAcc1)
	fmt.Println(lstAccount)

	myAcc2 := App.Account{"Satish", "s", 3, 30000}
	lstAccount = App.AddAccount(myAcc2)
	fmt.Println(lstAccount)

	//Delete Accounts
	App.RemoveAccount(1)
	fmt.Println(lstAccount)

	//Balance
	lstAccounts := App.ListAccounts()
	for k, _ := range lstAccounts {
		fmt.Printf("Balance for acct holder %d  : %d \n", lstAccounts[k].AccountNumber, App.AccountBalance(k))
	}

	//Deposit Money
	fmt.Printf("Balance before deposit: %d \n", lstAccounts[2].Balance)
	fmt.Printf("Balance after deposit: %d \n", App.DepositMoney(2, 500))

	//Withdraw Money
	fmt.Printf("Balance before deposit: %d \n", lstAccounts[3].Balance)
	fmt.Printf("Balance after deposit: %d \n", App.WithdrawMoney(3, 500))

	//Transfer Money
	fmt.Printf("Balance of the depositor after transfer: %d \n", lstAccounts[2].Balance)
	fmt.Printf("Balance of the depositee after transfer: %d \n", lstAccounts[3].Balance)
	depositorBalance, DepositeeBalance := App.TransferMoney(2, 3, 10000)
	fmt.Printf("Balance of the depositor after transfer: %d \n", depositorBalance)
	fmt.Printf("Balance of the depositee after transfer: %d \n", DepositeeBalance)
}
