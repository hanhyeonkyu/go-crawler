package main

import (
	"dev/go-crawler/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("alex")
	fmt.Println(account)
}
