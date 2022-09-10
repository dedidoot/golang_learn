package main

type CryptocurrencyFilter func(string) string

func getBlockchainFilter(name string, filter CryptocurrencyFilter) string {
	filteredName := filter(name)
	return filteredName
}

func cryptocurrencyFilter(name string) string {
	if name == "BTC" {
		return "Bitcoin"
	} else if name == "ETH" {
		return "Ethereum"
	} else if name == "ADA" {
		return "Cardano"
	} else if name == "SOL" {
		return "Solana"
	} else {
		return "Cryptocurrency not yet support in here..."
	}
}

func main() {
	println(getBlockchainFilter("ADA", cryptocurrencyFilter))
}
