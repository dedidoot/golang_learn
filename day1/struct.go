package main

import "fmt"

type Blockchain struct {
	Name, Currency, Date string
	PriceToUsd           int64
}

func main() {

	var blockchain Blockchain
	blockchain.Name = "Ethereum"
	blockchain.Currency = "ETH"
	blockchain.Date = "2022-09-14, 11:20:37 PM"
	blockchain.PriceToUsd = 1602.000000
	fmt.Println(blockchain)

}
