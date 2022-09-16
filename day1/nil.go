package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func blockchainFeature(feature string) map[string]string {
	if feature == "" {
		return nil
	} else {
		return map[string]string{
			"id":           strconv.Itoa(rand.Intn(99)),
			"feature_name": feature,
		}
	}
}

func main() {

	var blockchain1 map[string]string = blockchainFeature("Point of Sales")
	var blockchain2 map[string]string = blockchainFeature("Electric Token Transaction")
	var blockchain3 map[string]string = blockchainFeature("")

	fmt.Println(blockchain1)
	fmt.Println(blockchain2)
	if blockchain3 == nil {
		fmt.Println("Blockchain third is nil")
	} else {
		fmt.Println(blockchain3)
	}
}
