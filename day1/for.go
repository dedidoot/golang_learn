package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Index", counter)
		counter++
	}

	for newCounter := 0; newCounter <= 10; newCounter++ {
		fmt.Println("newCounter", newCounter)
	}

	blockchain := []string{"Bitcoin", "Ethereum", "Cardano", "Solana"}

	for index := 0; index < len(blockchain); index++ {
		fmt.Println("data", blockchain[index])
	}

	for item := range blockchain {
		fmt.Println("blockchain name", item)
	}

	frontEnd := make(map[string]string)
	frontEnd["javascript"] = "reactjs"
	frontEnd["typescript"] = "angular"

	for item := range frontEnd {
		parse := "Front End (" + item + ")"
		//fmt.Printf("Front End (%s)", item)
		println(parse)
	}
}
