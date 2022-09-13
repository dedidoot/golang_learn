package main

import "fmt"

func getRecover() {
	errorMessage := recover()
	if errorMessage != nil {
		fmt.Println("error message:", errorMessage)
	}
	println("Loaded...")
}

func runningApp(isError bool) {
	defer getRecover()
	if isError {
		panic("value error...")
	}

	println("Running...")
}

func main() {
	runningApp(true)
	println("done...")
}
