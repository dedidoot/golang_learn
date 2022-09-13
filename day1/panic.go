package main

func getLoaded() {
	println("Loaded...")
}

func runApplication(isError bool) {
	defer getLoaded()
	if isError {
		panic("error...")
	}

	println("Running...")
}

func main() {
	runApplication(true)
}
