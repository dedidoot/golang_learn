package main

import "fmt"

type Blockchain struct {
	Name, Currency, Date string
	PriceToUsd           float64
}

func main() {

	var blockchain Blockchain
	blockchain.Name = "Ethereum"
	blockchain.Currency = "ETH"
	blockchain.Date = "2022-09-14, 11:20:37 PM"
	blockchain.PriceToUsd = 1602.000000
	fmt.Println(blockchain)

	var btc = Blockchain{
		Name:       "Bitcoin",
		Currency:   "BTC",
		Date:       "2022-09-14, 11:53:37 PM",
		PriceToUsd: 20230.000000,
	}
	fmt.Println(btc)

}
