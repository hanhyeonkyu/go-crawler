package main

import (
	"dev/go-crawler/accounts"
	"dev/go-crawler/mydict"
	"fmt"
	"log"
)

func main() {
	account := accounts.NewAccount("alex")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("dexter")
	fmt.Println(account.Balance(), account.Owner())
	sampledict := mydict.SDictionary{"first": "First Word"}
	definite, error := sampledict.Search("first")
	if error != nil {
		fmt.Println(definite)
	} else {
		fmt.Println(err)
	}
}
