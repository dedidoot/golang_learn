package main

func getTwoBlockchainFavorite() (string, string, string) {
	return "Cardano", "Ethereum", "Bitcoin"
}

func main() {
	myFirstFavorite, mySecondFavorite, _ := getTwoBlockchainFavorite()

	println(myFirstFavorite)
	println(mySecondFavorite)
}
