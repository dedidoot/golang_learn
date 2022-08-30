package main

import "fmt"

func main() {
	var fullname string
	fullname = "John Doe"
	fmt.Println(fullname)

	fullname = "John Wick"
	fmt.Println(fullname)

	var birthday = 1988
	fmt.Println(birthday)

	city := "Yogyakarta"
	fmt.Println(city)

	codeArea := 2133
	fmt.Println(codeArea)

	var (
		firstName = "John"
		lastName  = "Doe"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
