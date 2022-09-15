package main

import "fmt"

type Blockchain struct {
	Name, Currency, Date string
	PriceToUsd           float64
}

func (blockchain Blockchain) getCrypto(name string) {
	fmt.Println("Hello", name)
}

func (blockchain Blockchain) getFavoriteCrypto(name string) {
	fmt.Println("My favorite crypto", name)
}

func main() {

	var blockchain Blockchain
	blockchain.Name = "Ethereum"
	blockchain.Currency = "ETH"
	blockchain.Date = "2022-09-14, 11:20:37 PM"
	blockchain.PriceToUsd = 1602.000000
	fmt.Println(blockchain)

	blockchain.getCrypto("Cardano")

	var btc = Blockchain{
		Name:       "Bitcoin",
		Currency:   "BTC",
		Date:       "2022-09-14, 11:53:37 PM",
		PriceToUsd: 20230.000000,
	}
	fmt.Println(btc)

	btc.getFavoriteCrypto("Ripple")
}
