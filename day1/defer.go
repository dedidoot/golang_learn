package main

func getLogging() {
	println("Good job!")
}

func getNews(page int) {
	defer getLogging() // always call whatever is error
	println("running news...")
	value := 10 / page
	println("page", value)
}

func main() {
	getNews(10)
}
