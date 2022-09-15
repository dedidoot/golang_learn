package main

import "fmt"

type SmartContract struct {
	Programming, Blockchain string
}

type Skills interface {
	GetSkill() (string, string)
}

func (contract SmartContract) GetSkill() (string, string) {
	return contract.Programming, contract.Blockchain
}

func WelcomeNewSkill(skill Skills) {
	fmt.Println(skill.GetSkill())
}

func main() {
	favoriteSkill := SmartContract{
		Programming: "Solidity",
		Blockchain:  "Ethereum 2.0",
	}

	WelcomeNewSkill(favoriteSkill)
}
