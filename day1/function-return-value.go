package main

func getBlockchain(name string) string {
	if name == "" {
		return "Please choose Blockchain"
	}
	return "Hello " + name
}

func main() {
	result := getBlockchain("Cardano")
	println(result)
}
