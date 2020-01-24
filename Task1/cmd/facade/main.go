package main

import (
	"fmt"
	"github.com/RyzhAlexWork/go-intern/Task1/pkg/facade"
	"github.com/RyzhAlexWork/go-intern/Task1/pkg/status"
	"github.com/RyzhAlexWork/go-intern/Task1/pkg/wallet"
)

func main() {
	newWallet := wallet.NewWallet(0, 1000000)
	newStatus := walletstatus.NewWalletStatus()
	newUser := facade.NewUser("Aksel", newWallet, newStatus)

	status := newUser.Add(500000)
	fmt.Printf("%s\n", status)
	status = newUser.Pay(200000)
	fmt.Printf("%s\n", status)
	balance := newUser.Balance()
	fmt.Printf("Balance: %d\n", balance)
}
