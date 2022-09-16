package main

import "fmt"

func getAnything(data interface{}) interface{} {
	return data
}

func main() {
	fmt.Println(getAnything("Hello"))
	fmt.Println(getAnything(1000))
	fmt.Println(getAnything(10.000))
	fmt.Println(getAnything(true))
}
