package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Rx"

	println(names[0])
	println(names[1])
	println(names[2])

	var rating = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(rating)
	fmt.Println(len(rating))
	fmt.Println(len(names))
}
