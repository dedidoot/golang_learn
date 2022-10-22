package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	myAddress1 := Address{
		City:     "Bantul",
		Province: "Sewon",
		Country:  "Indonesia",
	}

	myAddress2 := myAddress1

	myAddress2.Province = "Kasihan"

	fmt.Println(myAddress1)
	fmt.Println(myAddress2)

}
