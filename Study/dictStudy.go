package dictStudy

import (
	"dev/go-crawler/accounts"
	"dev/go-crawler/mydict"
	"fmt"
	"log"
)

func DictStudy1() {
	account := accounts.NewAccount("alex")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Println(err)
	}
	account.ChangeOwner("dexter")
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err1 := dictionary.Add(word, definition)
	if err1 != nil {
		fmt.Println(err1)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Update(word, "Second")
	if err2 != nil {
		fmt.Println(err2)
	}
	updated, _ := dictionary.Search(word)
	fmt.Println(updated)
	dictionary.Delete(word)
	_, err3 := dictionary.Search(word)
	if err3 != nil {
		fmt.Print(err3)
	}
}
