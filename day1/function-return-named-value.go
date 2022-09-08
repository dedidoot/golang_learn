package main

func getBlockchainFavorite() (favorite1, favorite2, favorite3 string, price int) {
	favorite1 = "Ethereum"
	favorite2 = "Cardano"
	favorite3 = "Solana"
	price = 9821331
	return
}

func main() {

	favorite1, favorite2, favorite3, price := getBlockchainFavorite()
	println(favorite1)
	println(favorite2)
	println(favorite3)
	println(price)
}
