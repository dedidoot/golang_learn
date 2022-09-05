package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John Doe",
		"address": "Las Vegas",
	}

	person["job"] = "Golang Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["job"])

	title := make(map[string]string)
	//title := map[string]string{}
	title["dr"] = "Doctor"
	title["MRes"] = "Master of Research"
	title["delete"] = "Deleted"

	delete(title, "delete")
	fmt.Println(title)
}
